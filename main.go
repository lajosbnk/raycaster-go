package main

import rl "github.com/gen2brain/raylib-go/raylib"

var player = Player{X: 0, Y: 0, Speed: 50}

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
	player.Update(rl.GetFrameTime())
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	player.Render()

	rl.EndDrawing()
}

func processInput() {

}
