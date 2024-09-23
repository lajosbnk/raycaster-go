package main

import (
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
	TurnSpeed:     270 * (math.Pi / 180),
	GameMap:       &gameMap,
}

func main() {
	initialize()

	for !rl.WindowShouldClose() {
		update()
		render()
	}

	rl.CloseWindow()
}

func initialize() {
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Raycaster")
	rl.SetWindowState(rl.FlagWindowUndecorated)
	rl.HideCursor()
	rl.SetTargetFPS(60)
}

func update() {
	player.Update()
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	gameMap.Render()
	player.Render()

	rl.EndDrawing()
}

func processInput() {

}
