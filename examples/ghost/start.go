package main

import "github.com/firefly-zero/firefly-go/firefly"

var startframes int

// logic for the start scene
func updateStart() {
	startframes += 1
	if startframes > 60 {
		buttons := firefly.ReadButtons(firefly.Combined)
		if buttons.N || buttons.S || buttons.E || buttons.W {
			scene = gamePlay
		}
	}
}

// render the start scene
func renderStart() {
	firefly.ClearScreen(firefly.ColorBlack)
	firefly.DrawText("GHOST CASTLE", titleFont, firefly.Point{X: 80, Y: 60}, firefly.ColorRed)
	firefly.DrawText("Press any button to Start", titleFont, firefly.Point{X: 44, Y: 100}, firefly.ColorRed)
}
