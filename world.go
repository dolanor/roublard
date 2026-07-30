package main

import (
	"github.com/bytearena/ecs"
	"github.com/g3n/engine/core"
)

var position *ecs.Component
var renderable *ecs.Component
var monster *ecs.Component
var health *ecs.Component
var meleeWeapon *ecs.Component
var armor *ecs.Component
var name *ecs.Component
var userMessage *ecs.Component

func InitializeWorld(startingLevel *Level) (*ecs.Manager, map[string]ecs.Tag, []core.INode) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()
	meshes := []core.INode{}

	player := manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable := manager.NewComponent()
	monster = manager.NewComponent()
	health = manager.NewComponent()
	meleeWeapon = manager.NewComponent()
	armor = manager.NewComponent()
	name = manager.NewComponent()
	userMessage = manager.NewComponent()

	playerMesh := loadElfMesh()
	meshes = append(meshes, playerMesh)
	playerMesh.SetVisible(true)

	// Get First Room
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()

	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{
			Image: playerMesh,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: x,
			Y: y,
			Z: playerMesh.Position().Z,
		}).
		AddComponent(health, &Health{
			MaxHealth:     30,
			CurrentHealth: 30,
		}).
		AddComponent(meleeWeapon, &MeleeWeapon{
			Name:          "Battle Axe",
			MinimumDamage: 10,
			MaximumDamage: 20,
			ToHitBonus:    3,
		}).
		AddComponent(armor, &Armor{
			Name:       "Plate Armor",
			Defense:    15,
			ArmorClass: 18,
		}).
		AddComponent(name, &Name{Label: "Player"}).
		AddComponent(userMessage, &UserMessage{
			AttackMessage:    "",
			DeadMessage:      "",
			GameStateMessage: "",
		})

	//Add a Monster in each room except the player's room
	for _, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			mX, mY := room.Center()

			//Flip a coin to see what to add...
			mobSpawn := GetDiceRoll(2)

			if mobSpawn == 1 {
				orcMesh := loadGoblinJanitorMesh()
				orcMesh.SetVisible(false)
				meshes = append(meshes, orcMesh)

				manager.NewEntity().
					AddComponent(monster, &Monster{}).
					AddComponent(renderable, &Renderable{
						Image: orcMesh,
					}).
					AddComponent(position, &Position{
						X: mX,
						Y: mY,
						Z: orcMesh.Position().Z,
					}).
					AddComponent(health, &Health{
						MaxHealth:     30,
						CurrentHealth: 30,
					}).
					AddComponent(meleeWeapon, &MeleeWeapon{
						Name:          "Broom of Doom",
						MinimumDamage: 4,
						MaximumDamage: 8,
						ToHitBonus:    1,
					}).
					AddComponent(armor, &Armor{
						Name:       "Leather boxer",
						Defense:    5,
						ArmorClass: 6,
					}).
					AddComponent(name, &Name{Label: "Goblin Janitor"}).
					AddComponent(userMessage, &UserMessage{
						AttackMessage:    "",
						DeadMessage:      "",
						GameStateMessage: "",
					})
			} else {
				skeletonMesh := loadSkeletonMesh()
				skeletonMesh.SetVisible(false)
				meshes = append(meshes, skeletonMesh)

				manager.NewEntity().
					AddComponent(monster, &Monster{}).
					AddComponent(renderable, &Renderable{
						Image: skeletonMesh,
					}).
					AddComponent(position, &Position{
						X: mX,
						Y: mY,
						Z: skeletonMesh.Position().Z,
					}).
					AddComponent(health, &Health{
						MaxHealth:     10,
						CurrentHealth: 10,
					}).
					AddComponent(meleeWeapon, &MeleeWeapon{
						Name:          "Short Sword",
						MinimumDamage: 2,
						MaximumDamage: 6,
						ToHitBonus:    0,
					}).
					AddComponent(armor, &Armor{
						Name:       "Bone",
						Defense:    3,
						ArmorClass: 4,
					}).
					AddComponent(name, &Name{Label: "Skeleton"}).
					AddComponent(userMessage, &UserMessage{
						AttackMessage:    "",
						DeadMessage:      "",
						GameStateMessage: "",
					})
			}

		}
	}

	players := ecs.BuildTag(player, position, health, meleeWeapon, armor, name, userMessage)
	tags["players"] = players

	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables

	monsters := ecs.BuildTag(monster, position, health, meleeWeapon, armor, name, userMessage)
	tags["monsters"] = monsters

	messengers := ecs.BuildTag(userMessage)
	tags["messengers"] = messengers

	return manager, tags, meshes
}
