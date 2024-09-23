package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	X             float32
	Y             float32
	Width         float32
	Height        float32
	RotationAngle float32
	WalkSpeed     float32
	TurnSpeed     float32
	TurnDirection int8
	WalkDirection int8
	GameMap       *Map
}

func (p *Player) Render() {
	rec := rl.Rectangle{
		X:      p.X * MINIMAP_SCALE_FACTOR,
		Y:      p.Y * MINIMAP_SCALE_FACTOR,
		Width:  p.Width * MINIMAP_SCALE_FACTOR,
		Height: p.Height * MINIMAP_SCALE_FACTOR,
	}
	rl.DrawRectangleRec(rec, rl.RayWhite)

	rl.DrawLine(
		int32(p.X*MINIMAP_SCALE_FACTOR),
		int32(p.Y*MINIMAP_SCALE_FACTOR),
		int32(p.X*MINIMAP_SCALE_FACTOR+float32(math.Cos(float64(p.RotationAngle))*40)),
		int32(p.Y*MINIMAP_SCALE_FACTOR+float32(math.Sin(float64(p.RotationAngle))*40)),
		rl.Red,
	)
}

func (p *Player) Update() {
	p.handleInput()

	p.RotationAngle += float32(p.TurnDirection) * (p.TurnSpeed * rl.GetFrameTime())
	moveStep := float32(p.WalkDirection) * p.WalkSpeed * rl.GetFrameTime()

	newPlayerX := p.X + float32(math.Cos(float64(p.RotationAngle)))*moveStep
	newPlayerY := p.Y + float32(math.Sin(float64(p.RotationAngle)))*moveStep

	if !gameMap.HasWallAt(newPlayerX, newPlayerY) {
		p.X = newPlayerX
		p.Y = newPlayerY
	}
}

func (p *Player) handleInput() {
	if rl.IsKeyDown(rl.KeyUp) {
		p.WalkDirection = 1
	}
	if rl.IsKeyDown(rl.KeyDown) {
		p.WalkDirection = -1
	}
	if rl.IsKeyDown(rl.KeyRight) {
		p.TurnDirection = 1
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		p.TurnDirection = -1
	}

	if rl.IsKeyReleased(rl.KeyUp) {
		p.WalkDirection = 0
	}
	if rl.IsKeyReleased(rl.KeyDown) {
		p.WalkDirection = 0
	}
	if rl.IsKeyReleased(rl.KeyRight) {
		p.TurnDirection = 0
	}
	if rl.IsKeyReleased(rl.KeyLeft) {
		p.TurnDirection = 0
	}
}
