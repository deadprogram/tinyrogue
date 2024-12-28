package main

import "github.com/firefly-zero/firefly-go/firefly"

// logic for the start scene
func updateStart() {
	pause += 1
	if pause > 60 {
		buttons := firefly.ReadButtons(firefly.Combined)
		if buttons.N || buttons.S || buttons.E || buttons.W {
			scene = gamePlay
			pause = 0
		}
	}
}

// render the start scene
func renderStart() {
	firefly.ClearScreen(firefly.ColorBlack)
	firefly.DrawText("GHOST CASTLE", titleFont, firefly.Point{X: 80, Y: 60}, firefly.ColorRed)
	firefly.DrawText("Press any button to Start", titleFont, firefly.Point{X: 44, Y: 100}, firefly.ColorRed)
}
