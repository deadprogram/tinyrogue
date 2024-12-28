package main

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

func updateGameover() {
	buttons := firefly.ReadButtons(firefly.Combined)
	if buttons.N || buttons.S || buttons.E || buttons.W {
		scene = gameStart
	}
}

func renderGameover() {
	firefly.ClearScreen(firefly.ColorBlack)

	firefly.DrawText("GAME OVER", titleFont, firefly.Point{X: 90, Y: 60}, firefly.ColorRed)
}
