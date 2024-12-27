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
	player := tinyrogue.NewPlayer("Player", &playerImage, 5)

	monsterImage := firefly.LoadFile("monster", nil).Image()
	monster := tinyrogue.NewCreature("Monster", &monsterImage, 60)
	monster.SetBehavior(tinyrogue.CreatureApproach)

	game.SetData(tinyrogue.NewGameData(16, 10, 16, 16))
	game.SetMap(tinyrogue.NewGameMap())

	game.SetPlayer(player)
	game.AddCreature(monster)

	game.SetActionSystem(&CombatSystem{})

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

type CombatSystem struct {
}

func (ca *CombatSystem) Action(attacker tinyrogue.Character, defender tinyrogue.Character) {
	firefly.LogDebug(attacker.Name() + " is attacking " + defender.Name())
}
