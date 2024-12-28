package main

import (
	"github.com/deadprogram/tinyrogue"
	"github.com/firefly-zero/firefly-go/firefly"
)

type Adventurer struct {
	*tinyrogue.Player
	*combatant
}

func NewAdventurer(name string, image *firefly.Image, speed int) *Adventurer {
	return &Adventurer{
		Player:    tinyrogue.NewPlayer(name, "adventurer", image, speed),
		combatant: NewCombatant(8, 10, 8, "sword"),
	}
}

type Ghost struct {
	*tinyrogue.Creature
	*combatant
}

func NewGhost(name string, image *firefly.Image, speed int) *Ghost {
	return &Ghost{
		Creature:  tinyrogue.NewCreature(name, "ghost", image, speed),
		combatant: NewCombatant(2, 5, 4, "shriek"),
	}
}
