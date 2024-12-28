package main

import (
	"github.com/deadprogram/tinyrogue"
	"github.com/firefly-zero/firefly-go/firefly"
)

type Adventurer struct {
	*tinyrogue.Player
}

func NewAdventurer(name string, image *firefly.Image, speed int) *Adventurer {
	return &Adventurer{
		Player: tinyrogue.NewPlayer(name, image, speed),
	}
}

type Ghost struct {
	*tinyrogue.Creature
}

func NewGhost(name string, image *firefly.Image, speed int) *Ghost {
	return &Ghost{
		Creature: tinyrogue.NewCreature(name, image, speed),
	}
}
