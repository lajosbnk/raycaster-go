package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	X     int32
	Y     int32
	Speed float32
}

func (p *Player) Render() {
	rl.DrawRectangle(player.X, player.Y, 20, 20, rl.Yellow)
}

func (p *Player) Update(dt float32) {
	p.X += int32(p.Speed * dt)
	p.Y += int32(p.Speed * dt)
}
