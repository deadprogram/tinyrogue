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

	game.LoadImage("forest")
	game.LoadImage("forest2")
	game.LoadImage("tree")
	game.LoadImage("tree2")

	gd := tinyrogue.NewGameData(16, 10, 16, 16)
	gd.FloorTypes = "forest,forest2"
	gd.WallTypes = "tree,tree2"
	gd.MinSize = 3
	gd.MaxSize = 6
	gd.MaxRooms = 36

	game.SetData(gd)
	game.SetMap(tinyrogue.NewGameMap())

	playerImage := game.LoadImage("player")
	player := tinyrogue.NewPlayer("Player", "player", playerImage, 5)

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
