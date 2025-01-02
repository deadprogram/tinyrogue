package main

import (
	"strconv"

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

	game.LoadImages("forest", "forest2", "tree", "tree2")

	gd := tinyrogue.NewGameData(16, 10, 16, 16)
	gd.MinSize = 3
	gd.MaxSize = 6
	gd.MaxRooms = 36
	game.SetData(gd)

	floors := []string{"forest", "forest2"}
	walls := []string{"tree", "tree2"}

	// create a few dungeons
	name := "Forest-"
	dungeons := make([]tinyrogue.Dungeon, 0)
	for i := 0; i < 4; i++ {
		d := tinyrogue.NewDungeon(name+"-"+strconv.Itoa(i), floors[i%2], walls[i%2])
		for j := 0; j < 7; j++ {
			nextLevel := tinyrogue.NewLevel(d.Name+"-"+strconv.Itoa(j), d.FloorTypes, d.WallTypes)
			d.Levels = append(d.Levels, nextLevel)
		}
		dungeons = append(dungeons, d)
	}

	// generate the first level of the first dungeon
	firstDungeon := &dungeons[0]
	firstLevel := firstDungeon.Levels[0]
	firstLevel.Generate()

	// put exit on first level to second level
	p := tinyrogue.NewPortal("portal", game.LoadImage("portal"), firstDungeon, firstDungeon.NextLevel(firstLevel))
	firstLevel.SetExit(p, firstLevel.OpenLocation())

	gm := tinyrogue.NewGameMap("Big World", dungeons, firstDungeon.Name, firstLevel.Name)
	game.SetMap(gm)

	player := tinyrogue.NewPlayer("Player", "player", game.LoadImage("player"), 5)
	player.ViewRadius = 3

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
