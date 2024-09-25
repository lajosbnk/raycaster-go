package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	MAP = [MAP_NUM_ROWS][MAP_NUM_COLS]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
)

type Map struct{}

func (m *Map) Render() {
	for i := 0; i < MAP_NUM_ROWS; i++ {
		for j := 0; j < MAP_NUM_COLS; j++ {
			tileX := j * TILE_SIZE
			tileY := i * TILE_SIZE
			var tileColor color.RGBA
			if MAP[i][j] == 0 {
				tileColor = color.RGBA{0, 0, 0, 255}
			} else {
				tileColor = color.RGBA{255, 255, 255, 255}
			}

			rec := rl.Rectangle{
				X:      float32(tileX) * MINIMAP_SCALE_FACTOR,
				Y:      float32(tileY) * MINIMAP_SCALE_FACTOR,
				Width:  float32(TILE_SIZE) * MINIMAP_SCALE_FACTOR,
				Height: float32(TILE_SIZE) * MINIMAP_SCALE_FACTOR,
			}
			rl.DrawRectangleRec(rec, tileColor)
		}
	}
}

func (m *Map) HasWallAt(x int, y int) bool {
	if x < 0 || x > WINDOW_WIDTH || y < 0 || y > WINDOW_HEIGHT {
		return true
	}

	mapGridIndexX := x / TILE_SIZE
	mapGridIndexY := y / TILE_SIZE

	if mapGridIndexX >= MAP_NUM_COLS || mapGridIndexY >= MAP_NUM_ROWS {
		return true
	}

	return MAP[mapGridIndexY][mapGridIndexX] != 0
}

func (m *Map) GetContentAt(x int, y int) int {
	if y < 0 || y >= MAP_NUM_ROWS || x < 0 || x >= MAP_NUM_COLS {
		return 0
	}
	return MAP[y][x]
}
