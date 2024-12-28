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

	player *Adventurer

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
				ghost := createGhost()

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
	game.UseFOV = true
	game.SetActionSystem(&CombatSystem{})

	loadGameImages()

	gd := tinyrogue.NewGameData(16, 10, 16, 16)
	gd.MinSize = 3
	gd.MaxSize = 6
	gd.MaxRooms = 32
	game.SetData(gd)

	startGame()
}

func startGame() {
	game.SetMap(tinyrogue.NewGameMap())

	createPlayer()
	ghost := createGhost()

	// set player initial position
	player.MoveTo(findSpawnLocation())

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

func loadGameImages() {
	floorImage := firefly.LoadFile("floor", nil).Image()
	game.Images["floor"] = &floorImage

	wallImage := firefly.LoadFile("wall", nil).Image()
	game.Images["wall"] = &wallImage

	playerImage := firefly.LoadFile("player", nil).Image()
	game.Images["player"] = &playerImage

	ghostImage := firefly.LoadFile("ghost", nil).Image()
	game.Images["ghost"] = &ghostImage
}

func createPlayer() {
	player = NewAdventurer("Sir Scaredy", game.Images["player"], 5)
	player.ViewRadius = 2
	game.SetPlayer(player)
}

func createGhost() *Ghost {
	ghost := NewGhost("Ghost", game.Images["ghost"], 60)
	ghost.SetBehavior(tinyrogue.CreatureApproach)
	game.AddCreature(ghost)

	return ghost
}

func removeAllGhosts() {
	for _, c := range game.Creatures {
		if gh, ok := c.(*Ghost); ok {
			removeGhost(gh)
		}
	}
}

func removeGhost(gh *Ghost) {
	game.RemoveCreature(gh)
	level := game.Map.CurrentLevel
	creaturePos := gh.GetPosition()
	level.Block(creaturePos.X, creaturePos.Y, false)
}
