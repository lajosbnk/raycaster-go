package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ray struct {
	RayAngle         float32
	WallHitX         float32
	WallHitY         float32
	Distance         float32
	WallHitContent   int
	WasHitVertical   bool
	IsRayFacingUp    bool
	IsRayFacingDown  bool
	IsRayFacingLeft  bool
	IsRayFacingRight bool
}

type Rays struct {
	Rays    []*Ray
	Player  *Player
	GameMap *Map
}

func (rs *Rays) castAllRays() {
	rayAngle := rs.Player.RotationAngle - (FOV_ANGLE / 2)

	for i := range rs.Rays {
		_castRay(rayAngle, i, rs)
		rayAngle += FOV_ANGLE / NUM_RAYS
	}
}

func (rs *Rays) Render() {
	for i := range rs.Rays {
		ray := rs.Rays[i]
		rl.DrawLine(
			int32(rs.Player.X*MINIMAP_SCALE_FACTOR),
			int32(rs.Player.Y*MINIMAP_SCALE_FACTOR),
			int32(ray.WallHitX*MINIMAP_SCALE_FACTOR),
			int32(ray.WallHitY*MINIMAP_SCALE_FACTOR),
			rl.Red)
	}
}

func _castRay(rayAngle float32, stripId int, rs *Rays) {
	rayAngle = _normalizeAngle(rayAngle)

	isRayFacingDown := rayAngle > 0 && rayAngle < math.Pi
	isRayFacingUp := !isRayFacingDown
	isRayFacingRight := (rayAngle < (math.Pi / 2)) || (rayAngle > (math.Pi * 1.5))
	isRayFacingLeft := !isRayFacingRight

	var xintercept, yintercept float32
	var xstep, ystep float32

	///////////////////////////////////////////
	// HORIZONTAL RAY GRID INTERSECTION CODE //
	///////////////////////////////////////////

	foundHorzWallHit := false
	horzWallHitX := 0
	horzWallHitY := 0
	horzWallContent := 0

	// Find the y-coordinate of the closest horizontal grid intersection
	yintercept = float32(math.Floor(float64(rs.Player.Y/float32(TILE_SIZE))) * float64(TILE_SIZE))
	if isRayFacingDown {
		yintercept += TILE_SIZE
	}

	// Find the x-coordinate of the closest horizontal grid intersection
	xintercept = rs.Player.X + (yintercept-rs.Player.Y)/float32(math.Tan(float64(rayAngle)))

	// Calculate the increment for xstep and ystep
	ystep = TILE_SIZE
	if isRayFacingUp {
		ystep *= -1
	}

	xstep = float32(TILE_SIZE) / float32(math.Tan(float64(rayAngle)))
	if isRayFacingLeft && xstep > 0 {
		xstep *= -1
	}
	if isRayFacingRight && xstep < 0 {
		xstep *= -1
	}

	nextHorzTouchX := xintercept
	nextHorzTouchY := yintercept

	// Increment xstep and ystep until we find a wall
	for nextHorzTouchX >= 0 && nextHorzTouchX <= WINDOW_WIDTH && nextHorzTouchY >= 0 && nextHorzTouchY <= WINDOW_HEIGHT {
		xToCheck := int(math.Floor(float64(nextHorzTouchX)))
		yToCheck := int(math.Floor(float64(nextHorzTouchY)))
		if isRayFacingUp {
			yToCheck += -1
		}

		if rs.GameMap.HasWallAt(xToCheck, yToCheck) {
			foundHorzWallHit = true
			horzWallHitX = int(nextHorzTouchX)
			horzWallHitY = int(nextHorzTouchY)

			contentXPos := int(math.Floor(float64(xToCheck) / TILE_SIZE))
			contentYPos := int(math.Floor(float64(yToCheck) / TILE_SIZE))
			horzWallContent = rs.GameMap.GetContentAt(contentXPos, contentYPos)

			break
		} else {
			nextHorzTouchX += xstep
			nextHorzTouchY += ystep
		}
	}

	///////////////////////////////////////////
	// VERTICAL RAY GRID INTERSECTION CODE //
	///////////////////////////////////////////

	foundVertWallHit := false
	vertWallHitX := 0
	vertWallHitY := 0
	vertWallContent := 0

	// Find the x-coordinate of the closest horizontal grid intersection
	xintercept = float32(math.Floor(float64(rs.Player.X/float32(TILE_SIZE))) * float64(TILE_SIZE))
	if isRayFacingRight {
		xintercept += TILE_SIZE
	}

	// Find the y-coordinate of the closest horizontal grid intersection
	yintercept = rs.Player.Y + (xintercept-rs.Player.X)*float32(math.Tan(float64(rayAngle)))

	// Calculate the increment for xstep and ystep
	xstep = TILE_SIZE
	if isRayFacingLeft {
		xstep *= -1
	}

	ystep = float32(TILE_SIZE) * float32(math.Tan(float64(rayAngle)))
	if isRayFacingUp && ystep > 0 {
		ystep *= -1
	}
	if isRayFacingDown && ystep < 0 {
		ystep *= -1
	}

	nextVertTouchX := xintercept
	nextVertTouchY := yintercept

	// Increment xstep and ystep until we find a wall
	for nextVertTouchX >= 0 && nextVertTouchX <= WINDOW_WIDTH && nextVertTouchY >= 0 && nextVertTouchY <= WINDOW_HEIGHT {
		xToCheck := int(math.Floor(float64(nextVertTouchX)))
		if isRayFacingLeft {
			xToCheck -= 1
		}
		yToCheck := int(math.Floor(float64(nextVertTouchY)))

		if rs.GameMap.HasWallAt(xToCheck, yToCheck) {
			foundVertWallHit = true
			vertWallHitX = int(nextVertTouchX)
			vertWallHitY = int(nextVertTouchY)

			contentXPos := int(math.Floor(float64(xToCheck) / TILE_SIZE))
			contentYPos := int(math.Floor(float64(yToCheck) / TILE_SIZE))
			vertWallContent = rs.GameMap.GetContentAt(contentXPos, contentYPos)

			break
		} else {
			nextVertTouchX += xstep
			nextVertTouchY += ystep
		}
	}

	horzHitDistance := float32(math.MaxInt)
	if foundHorzWallHit {
		horzHitDistance = _distanceBetweenPoints(rs.Player.X, rs.Player.Y, float32(horzWallHitX), float32(horzWallHitY))
	}

	vertHitDistance := float32(math.MaxInt)
	if foundVertWallHit {
		vertHitDistance = _distanceBetweenPoints(rs.Player.X, rs.Player.Y, float32(vertWallHitX), float32(vertWallHitY))
	}

	if vertHitDistance < horzHitDistance {
		rs.Rays[stripId].Distance = vertHitDistance
		rs.Rays[stripId].WallHitX = float32(vertWallHitX)
		rs.Rays[stripId].WallHitY = float32(vertWallHitY)
		rs.Rays[stripId].WallHitContent = vertWallContent
		rs.Rays[stripId].WasHitVertical = true
	} else {
		rs.Rays[stripId].Distance = horzHitDistance
		rs.Rays[stripId].WallHitX = float32(horzWallHitX)
		rs.Rays[stripId].WallHitY = float32(horzWallHitY)
		rs.Rays[stripId].WallHitContent = horzWallContent
		rs.Rays[stripId].WasHitVertical = false
	}

	rs.Rays[stripId].RayAngle = rayAngle
	rs.Rays[stripId].IsRayFacingDown = isRayFacingDown
	rs.Rays[stripId].IsRayFacingUp = isRayFacingUp
	rs.Rays[stripId].IsRayFacingLeft = isRayFacingLeft
	rs.Rays[stripId].IsRayFacingRight = isRayFacingRight
}

func _distanceBetweenPoints(x1 float32, y1 float32, x2 float32, y2 float32) float32 {
	xSquared := float64((x2 - x1) * (x2 - x1))
	ySquared := float64((y2 - y1) * (y2 - y1))

	return float32(math.Sqrt(xSquared + ySquared))
}

func _normalizeAngle(angle float32) float32 {
	angle = float32(math.Remainder(float64(angle), math.Pi*2))
	if angle < 0 {
		angle = math.Pi*2 + angle
	}

	return angle
}
