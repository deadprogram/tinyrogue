package main

import (
	"github.com/deadprogram/tinyrogue"
	"github.com/firefly-zero/firefly-go/firefly"
)

type CombatSystem struct {
}

func (ca *CombatSystem) Action(attacker tinyrogue.Character, defender tinyrogue.Character) {
	firefly.LogDebug(attacker.Name() + " is attacking " + defender.Name())
}
