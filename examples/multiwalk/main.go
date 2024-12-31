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
	game.UseFOV = true

	game.LoadImage("forest")
	game.LoadImage("forest2")
	game.LoadImage("tree")
	game.LoadImage("tree2")
	game.LoadImage("portal")

	gd := tinyrogue.NewGameData(16, 10, 16, 16)
	gd.MinSize = 3
	gd.MaxSize = 6
	gd.MaxRooms = 36
	game.SetData(gd)

	floors := "forest,forest2"
	walls := "tree,tree2"
	gm := tinyrogue.NewGeneratedGameMap("Big Forest", 1, 5, floors, walls)
	game.SetMap(gm)

	playerImage := game.LoadImage("player")
	player := tinyrogue.NewPlayer("Player", "player", playerImage, 5)
	player.ViewRadius = 34

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