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
		c := tinyrogue.CurrentGame().GetCreatureByName(attacker.Name())
		gh := c.(*Ghost)
		attackerWeaponClass = gh.WeaponClass()
		attackerWeaponName = gh.WeaponName()
	default:
		firefly.LogDebug("Unknown attacker kind: " + attacker.Kind())
	}

	switch defender.Kind() {
	case "adventurer":
		defenderArmorClass = player.ArmorClass()
	case "ghost":
		gh := defender.(*Ghost)
		defenderArmorClass = gh.ArmorClass()
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
			gh := defender.(*Ghost)
			remainingHealth := gh.Damage(damageRoll)
			if remainingHealth <= 0 {
				// Ghost defeated!
				msg2 = "Critical hit for " + strconv.Itoa(damageRoll) + " damage! " + gh.Name() + " defeated!"
				firefly.LogDebug(msg2)

				// Remove ghost from the game
				removeGhost(gh)
				score++

				// if all ghosts are defeated, respawn a new batch
				if len(game.Creatures) == 0 {
					respawnGhost = true
					numberGhosts++
				}
			}
		}

		firefly.LogDebug(msg1 + " " + msg2)
		dialog := tinyrogue.NewDialog(msg1, msg2, &msgFont, firefly.ColorRed, firefly.ColorBlack, true)
		tinyrogue.CurrentGame().ShowDialog(dialog)
	} else {
		msg1 := attacker.Name() + " tries " + attackerWeaponName + " on " + defender.Name()
		msg2 := "but it misses."
		firefly.LogDebug(msg1 + " " + msg2)

		dialog := tinyrogue.NewDialog(msg1, msg2, &msgFont, firefly.ColorRed, firefly.ColorBlack, true)
		tinyrogue.CurrentGame().ShowDialog(dialog)
	}
}
