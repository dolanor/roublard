package main

import (
	"fmt"

	"github.com/g3n/engine/gui"
)

var hudLabel *gui.Label

func CreateHUDWindow(panel *gui.Panel, width, height int) {

	w1 := gui.NewWindow(float32(width), float32(height))
	w1.SetPosition(float32(width), 3*float32(height))
	w1.SetResizable(true)

	hudLabel = gui.NewLabel("placeholder")

	w1.Add(hudLabel)
	panel.Add(w1)
}

func ProcessHUDG3N(g *Game) {
	for _, p := range g.World.Query(g.WorldTags["players"]) {
		h := p.Components[healths].(*Health)
		healthText := fmt.Sprintf("Health: %d / %d", h.CurrentHealth, h.MaxHealth)

		ac := p.Components[armors].(*Armor)
		acText := fmt.Sprintf("Armor Class: %d", ac.ArmorClass)

		defText := fmt.Sprintf("Defense: %d", ac.Defense)

		wpn := p.Components[meleeWeapons].(*MeleeWeapon)
		dmg := fmt.Sprintf("Damage: %d - %d", wpn.MinimumDamage, wpn.MaximumDamage)
		bonus := fmt.Sprintf("To Hit Bonus: %d", wpn.ToHitBonus)

		hudLabel.SetText(fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", healthText, acText, defText, dmg, bonus))
	}
}
