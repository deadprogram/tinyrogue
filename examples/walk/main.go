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

	game.LoadImages("forest", "forest2", "tree", "tree2")

	gd := tinyrogue.NewGameData(16, 10, 16, 16)
	gd.MinSize = 3
	gd.MaxSize = 6
	gd.MaxRooms = 36
	game.SetData(gd)

	floors := "forest,forest2"
	walls := "tree,tree2"
	game.SetMap(tinyrogue.NewSingleGameMapWithTerrain(floors, walls))

	player := tinyrogue.NewPlayer("Player", "player", game.LoadImage("player"), 5)

	game.SetPlayer(player)

	// set player initial position
	player.MoveTo(game.CurrentLevel().OpenLocation())
}

func update() {
	game.Update()
}

func render() {
	game.Render()
}
