package main

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var gameMap = Map{}
var player = Player{
	X:             WINDOW_WIDTH / 2,
	Y:             WINDOW_HEIGHT / 2,
	Width:         5,
	Height:        5,
	TurnDirection: 0,
	WalkDirection: 0,
	RotationAngle: math.Pi / 2,
	WalkSpeed:     200,
	TurnSpeed:     90 * (math.Pi / 180),
	GameMap:       &gameMap,
}
var rays = Rays{Player: &player, GameMap: &gameMap, Rays: make([]*Ray, NUM_RAYS)}

var colorBuffer = make([]color.RGBA, WINDOW_WIDTH*WINDOW_HEIGHT)
var texture rl.Texture2D

func main() {
	initialize()

	for !rl.WindowShouldClose() {
		update()
		render()
	}

	rl.UnloadTexture(texture)
	rl.CloseWindow()
}

func initialize() {
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Raycaster")
	rl.SetWindowState(rl.FlagWindowUndecorated)
	rl.HideCursor()
	rl.SetTargetFPS(60)

	for i := range NUM_RAYS {
		rays.Rays[i] = &Ray{}
	}

	texture = rl.LoadTextureFromImage(rl.GenImageColor(WINDOW_WIDTH, WINDOW_HEIGHT, rl.Black))
}

func update() {
	player.Update()
	rays.castAllRays()
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	generate3DProjection()

	rl.UpdateTexture(texture, colorBuffer)
	rl.DrawTexture(texture, 0, 0, rl.White)
	clearColorBuffer(color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF})

	gameMap.Render()
	rays.Render()
	player.Render()

	rl.EndDrawing()

}

func generate3DProjection() {
	for i := 0; i < NUM_RAYS; i++ {
		perpDistance := rays.Rays[i].Distance * float32(math.Cos(float64(rays.Rays[i].RayAngle)-float64(player.RotationAngle)))
		distanceProjPlane := (WINDOW_WIDTH / 2) / math.Tan(FOV_ANGLE/2)
		projectedWallHeight := (TILE_SIZE / perpDistance) * float32(distanceProjPlane)
		wallStripHeight := int(projectedWallHeight)

		wallTopPixel := (WINDOW_HEIGHT / 2) - (wallStripHeight / 2)
		if wallTopPixel < 0 {
			wallTopPixel = 0
		}

		wallBottomPixel := (WINDOW_HEIGHT / 2) + (wallStripHeight / 2)
		if wallBottomPixel > WINDOW_HEIGHT {
			wallBottomPixel = WINDOW_HEIGHT
		}

		for y := wallTopPixel; y < wallBottomPixel; y++ {
			if rays.Rays[i].WasHitVertical {
				colorBuffer[WINDOW_WIDTH*y+i] = color.RGBA{R: 0xCC, G: 0xCC, B: 0xCC, A: 0xFF}
			} else {
				colorBuffer[WINDOW_WIDTH*y+i] = color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
			}
		}
	}
}

func clearColorBuffer(pixelColor color.RGBA) {
	for x := 0; x < WINDOW_WIDTH; x++ {
		for y := 0; y < WINDOW_HEIGHT; y++ {
			colorBuffer[(WINDOW_WIDTH*y)+x] = pixelColor
		}
	}
}
