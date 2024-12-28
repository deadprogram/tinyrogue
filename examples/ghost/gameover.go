package main

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

func updateGameover() {
	pause++
	if pause > 180 {
		scene = gameStart
		pause = 0
	}
}

func renderGameover() {
	firefly.ClearScreen(firefly.ColorBlack)

	firefly.DrawText("GAME OVER", titleFont, firefly.Point{X: 90, Y: 60}, firefly.ColorRed)
}
