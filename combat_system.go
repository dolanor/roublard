package main

import (
	"fmt"

	"github.com/bytearena/ecs"
)

func AttackSystem(g *Game, attackerPosition *Position, defenderPosition *Position) {
	var attacker *ecs.QueryResult = nil
	var defender *ecs.QueryResult = nil

	//Get the attacker and defender if either is a player
	for _, playerCombatant := range g.World.Query(g.WorldTags["players"]) {
		pos := playerCombatant.Components[positions].(*Position)

		if pos.IsEqual(attackerPosition) {
			//This is the attacker
			attacker = playerCombatant
		} else if pos.IsEqual(defenderPosition) {
			//This is the defender
			defender = playerCombatant
		}
	}

	//Get the attacker and defender if either is a monster
	for _, cbt := range g.World.Query(g.WorldTags["monsters"]) {
		pos := cbt.Components[positions].(*Position)

		if pos.IsEqual(attackerPosition) {
			//This is the attacker
			attacker = cbt
		} else if pos.IsEqual(defenderPosition) {
			//This is the defender
			defender = cbt
		}

	}
	//If we somehow don't have an attacker or defender, just leave
	if attacker == nil || defender == nil {
		return
	}
	//Grab the required information
	defenderArmor := defender.Components[armors].(*Armor)
	defenderHealth := defender.Components[healths].(*Health)
	defenderName := defender.Components[names].(*Name).Label
	defenderMessage := defender.Components[userMessages].(*UserMessage)

	attackerWeapon := attacker.Components[meleeWeapons].(*MeleeWeapon)
	attackerName := attacker.Components[names].(*Name).Label
	attackerMessage := attacker.Components[userMessages].(*UserMessage)

	//if the attacker is dead, don't let them attackerWeapon
	if attacker.Components[healths].(*Health).CurrentHealth <= 0 {
		return
	}
	//Roll a d10 to hit
	toHitRoll := GetDiceRoll(10)

	if toHitRoll+attackerWeapon.ToHitBonus > defenderArmor.ArmorClass {
		// It's a hit!
		damageRoll := GetRandomBetween(attackerWeapon.MinimumDamage, attackerWeapon.MaximumDamage)

		damageDone := damageRoll - defenderArmor.Defense
		// Let's not have the weapon heal the defender
		if damageDone < 0 {
			damageDone = 0
		}
		defenderHealth.CurrentHealth -= damageDone
		attackerMessage.AttackMessage = fmt.Sprintf("%s swings %s at %s and hits for %d health.", attackerName, attackerWeapon.Name, defenderName, damageDone)

		if defenderHealth.CurrentHealth <= 0 {
			r, ok := defender.Components[renderables].(*Renderable)
			if ok {
				go animateDeath(r)
			}

			defenderMessage.DeadMessage = fmt.Sprintf("%s has died!", defenderName)
			if defenderName == "Player" {
				defenderMessage.GameStateMessage = "Game Over!"
				g.Turn = GameOver
			}
		}

	} else {
		attackerMessage.AttackMessage = fmt.Sprintf("%s swings %s at %s and misses.", attackerName, attackerWeapon.Name, defenderName)
	}
}
