package main

import (
	"strconv"

	"github.com/deadprogram/tinyrogue"
	"github.com/firefly-zero/firefly-go/firefly"
)

type Combatant interface {
	ArmorClass() int
	Health() int
	WeaponClass() int
	WeaponName() string
	Damage(int) int
}

type combatant struct {
	armorClass  int
	health      int
	weaponClass int
	weaponName  string
}

func NewCombatant(armorClass, health, weaponClass int, weaponName string) *combatant {
	return &combatant{
		armorClass:  armorClass,
		health:      health,
		weaponClass: weaponClass,
		weaponName:  weaponName,
	}
}

func (c *combatant) ArmorClass() int {
	return c.armorClass
}

func (c *combatant) Health() int {
	return c.health
}

func (c *combatant) WeaponClass() int {
	return c.weaponClass
}

func (c *combatant) WeaponName() string {
	return c.weaponName
}

func (c *combatant) Damage(damage int) int {
	c.health -= damage
	return c.health
}

type CombatSystem struct {
}

func (ca *CombatSystem) Action(attacker tinyrogue.Character, defender tinyrogue.Character) {
	firefly.LogDebug(attacker.Name() + " is attacking " + defender.Name())

	var attackerWeaponClass, defenderArmorClass int
	var attackerWeaponName string

	switch attacker.Kind() {
	case "adventurer":
		attackerWeaponClass = player.WeaponClass()
		attackerWeaponName = player.WeaponName()
	case "ghost":
		attackerWeaponClass = ghost.WeaponClass()
		attackerWeaponName = ghost.WeaponName()
	default:
		firefly.LogDebug("Unknown attacker kind: " + attacker.Kind())
	}

	switch defender.Kind() {
	case "adventurer":
		defenderArmorClass = player.ArmorClass()
	case "ghost":
		defenderArmorClass = ghost.ArmorClass()
	default:
		firefly.LogDebug("Unknown defender kind: " + defender.Kind())
	}

	// Roll a d20 to hit
	toHitRoll := tinyrogue.GetDiceRoll(20)
	if toHitRoll > defenderArmorClass {
		// It's a hit!
		damageRoll := tinyrogue.GetDiceRoll(attackerWeaponClass)

		msg1 := attacker.Name() + " uses " + attackerWeaponName + " on " + defender.Name()
		msg2 := "and hits for " + strconv.Itoa(damageRoll) + " damage!"

		// Apply damage
		switch defender.Kind() {
		case "adventurer":
			remainingHealth := player.Damage(damageRoll)
			if remainingHealth <= 0 {
				// We're dead!
				msg2 = strconv.Itoa(damageRoll) + " damage! You are dead!"
				firefly.LogDebug("Game over!")
				game.Turn = tinyrogue.GameOver
			}
		case "ghost":
			remainingHealth := ghost.Damage(damageRoll)
			if remainingHealth <= 0 {
				// Ghost defeated!
				msg2 = "Critical hit for " + strconv.Itoa(damageRoll) + " damage! Ghost defeated!"
				firefly.LogDebug("Ghost defeated!")

				// Remove ghost from the game
				game.RemoveCreature(ghost)
				level := game.Map.CurrentLevel
				creaturePos := ghost.GetPosition()
				level.Block(creaturePos.X, creaturePos.Y, false)

				respawnGhost = true
			}
		}

		firefly.LogDebug(msg1 + " " + msg2)
		dialog := tinyrogue.NewMessage(msg1, &msgFont, firefly.ColorRed, firefly.ColorBlack, true)
		dialog.Text2 = msg2
		tinyrogue.CurrentGame().ShowMessage(dialog)
	} else {
		msg := attacker.Name() + " tries " + attackerWeaponName + " on " + defender.Name() + " and misses."
		firefly.LogDebug(msg)
		dialog := tinyrogue.NewMessage(msg, &msgFont, firefly.ColorRed, firefly.ColorBlack, true)
		tinyrogue.CurrentGame().ShowMessage(dialog)
	}
}
