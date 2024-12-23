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

	floorImage := firefly.LoadFile("floor").Image()
	game.Images["floor"] = &floorImage
	wallImage := firefly.LoadFile("wall").Image()
	game.Images["wall"] = &wallImage

	game.SetData(tinyrogue.NewGameData(15, 15))
	game.SetMap(tinyrogue.NewGameMap())
}

func update() {
	game.Update()
}

func render() {
	game.Render()
}

func main() {
}
