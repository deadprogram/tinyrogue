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

	floorImage := firefly.LoadFile("floor", nil).Image()
	game.Images["floor"] = &floorImage

	wallImage := firefly.LoadFile("wall", nil).Image()
	game.Images["wall"] = &wallImage

	playerImage := firefly.LoadFile("player", nil).Image()
	player := tinyrogue.NewPlayer("Player", "player", &playerImage, 5)

	game.SetData(tinyrogue.NewGameData(16, 10, 16, 16))
	game.SetMap(tinyrogue.NewGameMap())

	game.SetPlayer(player)

	// set player initial position
	entrance := tinyrogue.Position{X: 2, Y: 2}
	player.MoveTo(entrance)
}

func update() {
	game.Update()
}

func render() {
	game.Render()
}
