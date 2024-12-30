package main

import (
	"github.com/deadprogram/tinyrogue"
	"github.com/firefly-zero/firefly-go/firefly"
)

func init() {
	firefly.Boot = boot
	firefly.Update = update
	firefly.Render = render
}

var game *tinyrogue.Game

func boot() {
	game = tinyrogue.NewGame()

	game.LoadImage("floor")
	game.LoadImage("wall")

	gd := tinyrogue.NewGameData(16, 10, 16, 16)

	game.SetData(gd)
	game.SetMap(tinyrogue.NewGameMap())

	player := tinyrogue.NewPlayer("Player", "player", game.LoadImage("player"), 5)
	game.SetPlayer(player)

	// set player initial position
	player.MoveTo(findOpenLocation())
}

func update() {
	game.Update()
}

func render() {
	game.Render()
}

func findOpenLocation() tinyrogue.Position {
	l := game.Map.CurrentLevel
	for {
		pos, free := l.RandomLocation()
		if free {
			return pos
		}
	}
}
