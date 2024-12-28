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
	pause = 0

	titleFont firefly.Font
	msgFont   firefly.Font

	game *tinyrogue.Game

	player       *Adventurer
	ghost        *Ghost
	respawnGhost bool
	respawnDelay int
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
		if game.MessageShowing {
			return
		}
		if game.Turn == tinyrogue.GameOver {
			scene = gameOver
			pause = 0
		}
		if respawnGhost {
			respawnDelay++
			if respawnDelay > 60 {
				createGhost()

				ghost.MoveTo(findSpawnLocation())
				respawnGhost = false
				respawnDelay = 0
			}
		}
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
	game.SetPlayer(player)

	ghostImage := firefly.LoadFile("ghost", nil).Image()
	game.Images["ghost"] = &ghostImage
	createGhost()

	gd := tinyrogue.NewGameData(16, 10, 16, 16)
	gd.MinSize = 3
	gd.MaxSize = 6
	gd.MaxRooms = 32
	game.SetData(gd)

	game.SetMap(tinyrogue.NewGameMap())
	game.UseFOV = true

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

func createGhost() {
	ghost = NewGhost("Ghost", game.Images["ghost"], 60)
	ghost.SetBehavior(tinyrogue.CreatureApproach)
	game.AddCreature(ghost)
}
