package main

import (
	"github.com/deadprogram/tinyrogue"
	"github.com/firefly-zero/firefly-go/firefly"
)

const (
	gameStart = "start"
	gamePlay  = "game"
	gameOver  = "gameover"
)

var (
	scene = gameStart

	titleFont firefly.Font
	msgFont   firefly.Font

	game *tinyrogue.Game

	player *Adventurer
	ghost  *Ghost
)

func init() {
	firefly.Boot = boot
	firefly.Update = update
	firefly.Render = render
}

func boot() {
	titleFont = firefly.LoadFile("titlefont", nil).Font()
	msgFont = firefly.LoadFile("msgfont", nil).Font()

	setupGame()
}

func update() {
	switch scene {
	case gameStart:
		updateStart()
	case gamePlay:
		game.Update()
	case gameOver:
		updateGameover()
	}
}

func render() {
	switch scene {
	case gameStart:
		renderStart()
	case gamePlay:
		game.Render()
	case gameOver:
		renderGameover()
	}
}

func setupGame() {
	game = tinyrogue.NewGame()

	floorImage := firefly.LoadFile("floor", nil).Image()
	game.Images["floor"] = &floorImage

	wallImage := firefly.LoadFile("wall", nil).Image()
	game.Images["wall"] = &wallImage

	playerImage := firefly.LoadFile("player", nil).Image()
	player = NewAdventurer("Sir Shaky", &playerImage, 5)
	player.ViewRadius = 2

	ghostImage := firefly.LoadFile("ghost", nil).Image()
	ghost = NewGhost("Ghost", &ghostImage, 60)
	ghost.SetBehavior(tinyrogue.CreatureApproach)

	gd := tinyrogue.NewGameData(16, 10, 16, 16)
	gd.MinSize = 3
	gd.MaxSize = 6
	gd.MaxRooms = 32
	game.SetData(gd)

	game.SetMap(tinyrogue.NewGameMap())
	game.UseFOV = true

	game.SetPlayer(player)
	game.AddCreature(ghost)

	game.SetActionSystem(&CombatSystem{})

	// set player initial position
	entrance := tinyrogue.Position{X: 1, Y: 2}
	player.MoveTo(entrance)

	// set monster initial position
	ghost.MoveTo(findSpawnLocation())
}

func findSpawnLocation() tinyrogue.Position {
	l := game.Map.CurrentLevel
	for {
		pos, free := l.RandomLocation()
		if free {
			return pos
		}
	}
}
