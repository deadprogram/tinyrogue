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
	// create a new game
	game = tinyrogue.NewGame()

	// load the image tiles for the floor and walls
	game.LoadImage("floor")
	game.LoadImage("wall")

	// set the dimensions for the game and the tiles
	gd := tinyrogue.NewGameData(16, 10, 16, 16)
	game.SetData(gd)

	// generate a random game map
	game.SetMap(tinyrogue.NewGameMap())

	// create the player
	player := tinyrogue.NewPlayer("Player", "player", game.LoadImage("player"), 5)
	game.SetPlayer(player)

	// set player initial position to some open spot on the map.
	player.MoveTo(game.CurrentLevel().OpenLocation())
}

func update() {
	game.Update()
}

func render() {
	game.Render()
}
