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
	player := tinyrogue.NewPlayer(&playerImage, 5)

	monsterImage := firefly.LoadFile("monster", nil).Image()
	monster := tinyrogue.NewCreature(&monsterImage, 60)
	monster.SetBehavior(tinyrogue.CreatureApproach)

	game.SetData(tinyrogue.NewGameData(16, 10))
	game.SetMap(tinyrogue.NewGameMap())

	game.SetPlayer(player)
	game.AddCreature(monster)

	// set player initial position
	entrance := tinyrogue.Position{X: 1, Y: 1}
	player.MoveTo(entrance)

	// set monster initial position
	monsterPos := tinyrogue.Position{X: 12, Y: 7}
	monster.MoveTo(monsterPos)
}

func update() {
	game.Update()
}

func render() {
	game.Render()
}
