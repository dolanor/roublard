package main

import (
	"github.com/bytearena/ecs"
	"github.com/g3n/engine/core"
)

var position *ecs.Component
var renderable *ecs.Component
var monster *ecs.Component

func InitializeWorld(startingLevel *Level, scene *core.Node) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	// WARNING: this is global state
	player := manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable := manager.NewComponent()
	monster = manager.NewComponent()

	elfMesh := loadElfMesh()
	scene.Add(elfMesh)
	elfMesh.SetVisible(true)

	// Get First Room
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()

	// Define the elf wizard in the ECS
	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{
			node: elfMesh,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: x,
			Y: y,
			Z: elfMesh.Position().Z,
		})

	addTorchLight(scene, manager, player, movable, x, y)

	//Add a Monster in each room except the player's room
	for _, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			monsterMesh := loadSkeletonMesh()
			monsterMesh.SetVisible(false)
			scene.Add(monsterMesh)

			mX, mY := room.Center()
			manager.NewEntity().
				AddComponent(monster, &Monster{
					Name: "Skeleton",
				}).
				AddComponent(renderable, &Renderable{
					node: monsterMesh,
				}).
				AddComponent(position, &Position{
					X: mX,
					Y: mY,
					Z: monsterMesh.Position().Z,
				})

		}
	}

	players := ecs.BuildTag(player, position)
	tags["players"] = players

	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables

	monsters := ecs.BuildTag(monster, position)
	tags["monsters"] = monsters

	return manager, tags
}
