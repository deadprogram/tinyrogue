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
	firefly.LogDebug("Boot")
	game = tinyrogue.NewGame()

	floorImage := firefly.LoadFile("floor", nil).Image()
	game.Images["floor"] = &floorImage

	wallImage := firefly.LoadFile("wall", nil).Image()
	game.Images["wall"] = &wallImage

	playerImage := firefly.LoadFile("player", nil).Image()
	game.Images["player"] = &playerImage
	player := tinyrogue.NewPlayer()
	player.SetImage(game.Images["player"])
	player.SetSpeed(10)

	monsterImage := firefly.LoadFile("monster", nil).Image()
	game.Images["monster"] = &monsterImage
	monster := tinyrogue.NewCreature()
	monster.SetImage(game.Images["monster"])
	monster.SetSpeed(60)

	firefly.LogDebug("set data")
	game.SetData(tinyrogue.NewGameData(16, 10))

	firefly.LogDebug("set map")
	game.SetMap(tinyrogue.NewGameMap())

	firefly.LogDebug("set player")
	game.SetPlayer(player)
	game.AddCreature(monster)

	// set player initial position
	entrance := &tinyrogue.Position{X: 1, Y: 1}
	player.MoveTo(entrance)

	// set monster initial position
	monsterPos := &tinyrogue.Position{X: 5, Y: 5}
	monster.MoveTo(monsterPos)
}

func update() {
	game.Update()
}

func render() {
	game.Render()
}

func main() {
}
