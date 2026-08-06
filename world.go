package main

import (
	"github.com/bytearena/ecs"
	"github.com/g3n/engine/core"
)

var positions *ecs.Component
var renderables *ecs.Component
var monsters *ecs.Component
var healths *ecs.Component
var meleeWeapons *ecs.Component
var armors *ecs.Component
var names *ecs.Component
var userMessages *ecs.Component

func InitializeWorldEntities(startingLevel *Level) (*ecs.Manager, map[string]ecs.Tag, []core.INode) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()
	meshes := []core.INode{}

	player := manager.NewComponent()
	positions = manager.NewComponent()
	renderables = manager.NewComponent()
	movable := manager.NewComponent()
	monsters = manager.NewComponent()
	healths = manager.NewComponent()
	meleeWeapons = manager.NewComponent()
	armors = manager.NewComponent()
	names = manager.NewComponent()
	userMessages = manager.NewComponent()

	playerMesh := loadElfMesh()
	meshes = append(meshes, playerMesh)
	playerMesh.SetVisible(true)

	// Get First Room
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()

	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderables, &Renderable{
			Image: playerMesh,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(positions, &Position{
			X: x,
			Y: y,
			Z: playerMesh.Position().Z,
		}).
		AddComponent(healths, &Health{
			Max:     30,
			Current: 30,
		}).
		AddComponent(meleeWeapons, &MeleeWeapon{
			Name:       "Battle Axe",
			MinDamage:  10,
			MaxDamage:  20,
			ToHitBonus: 3,
		}).
		AddComponent(armors, &Armor{
			Name:       "Plate Armor",
			Defense:    15,
			ArmorClass: 18,
		}).
		AddComponent(names, &Name{Label: "Player"}).
		AddComponent(userMessages, &UserMessage{
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
				goblinJanitorMesh := loadGoblinJanitorMesh()
				goblinJanitorMesh.SetVisible(false)
				meshes = append(meshes, goblinJanitorMesh)

				newMonster(manager, goblinJanitorMesh,
					mX, mY, goblinJanitorMesh.Position().Z,
					30, 30,
					&MeleeWeapon{
						Name:       "Broom of Doom",
						MinDamage:  4,
						MaxDamage:  8,
						ToHitBonus: 1,
					},
					&Armor{
						Name:       "Leather boxer",
						Defense:    5,
						ArmorClass: 6,
					},
					"Goblin Janitor",
				)
			} else {
				skeletonMesh := loadSkeletonMesh()
				skeletonMesh.SetVisible(false)
				meshes = append(meshes, skeletonMesh)

				newMonster(manager, skeletonMesh,
					mX, mY, skeletonMesh.Position().Z,
					10, 10,
					&MeleeWeapon{
						Name:       "Short Sword",
						MinDamage:  2,
						MaxDamage:  6,
						ToHitBonus: 0,
					},
					&Armor{
						Name:       "Bone",
						Defense:    3,
						ArmorClass: 4,
					},
					"Skeleton",
				)
			}
		}
	}

	playerTags := ecs.BuildTag(player, positions, healths, meleeWeapons, armors, names, userMessages)
	tags["players"] = playerTags

	renderableTags := ecs.BuildTag(renderables, positions)
	tags["renderables"] = renderableTags

	monsterTags := ecs.BuildTag(monsters, positions, healths, meleeWeapons, armors, names, userMessages, renderables)
	tags["monsters"] = monsterTags

	messengerTags := ecs.BuildTag(userMessages)
	tags["messengers"] = messengerTags

	return manager, tags, meshes
}

func newMonster(manager *ecs.Manager, mesh core.INode, x, y int, z float32, currentHealth, maxHealth int, meleeWeapon, armor any, name string) {
	manager.NewEntity().
		AddComponent(monsters, &Monster{}).
		AddComponent(renderables, &Renderable{
			Image: mesh,
		}).
		AddComponent(positions, &Position{
			X: x,
			Y: y,
			Z: z,
		}).
		AddComponent(healths, &Health{
			Current: currentHealth,
			Max:     maxHealth,
		}).
		AddComponent(meleeWeapons, meleeWeapon).
		AddComponent(armors, armor).
		AddComponent(names, &Name{Label: name}).
		AddComponent(userMessages, &UserMessage{})
}
