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

		firefly.LogDebug(attacker.Name() + " uses " + attackerWeaponName + " on " + defender.Name() + " and hits for " + strconv.Itoa(damageRoll) + " health.")
		msg := tinyrogue.NewMessage(attacker.Name()+" uses "+attackerWeaponName+" on "+defender.Name(),
			&msgFont, firefly.ColorRed, firefly.ColorBlack, true)
		msg.Text2 = "and hits for " + strconv.Itoa(damageRoll) + " damage!"
		tinyrogue.CurrentGame().ShowMessage(msg)
	} else {
		firefly.LogDebug(attacker.Name() + " tries " + attackerWeaponName + " on " + defender.Name() + " and misses.")
		msg := tinyrogue.NewMessage(attacker.Name()+" tries "+attackerWeaponName+" on "+defender.Name()+" and misses.",
			&msgFont, firefly.ColorRed, firefly.ColorBlack, true)
		tinyrogue.CurrentGame().ShowMessage(msg)
	}

	// TODO: Implement health reduction
}
