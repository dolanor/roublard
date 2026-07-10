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

	playerImg := loadElfMesh()
	scene.Add(playerImg)
	playerImg.SetVisible(true)

	// Get First Room
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()

	// Define the elf wizard in the ECS
	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{
			Image: playerImg,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: x,
			Y: y,
			Z: playerImg.Position().Z,
		})

	//Add a Monster in each room except the player's room
	for _, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			skellyImg := loadSkeletonMesh()
			skellyImg.SetVisible(false)
			scene.Add(skellyImg)

			mX, mY := room.Center()
			manager.NewEntity().
				AddComponent(monster, &Monster{
					Name: "Skeleton",
				}).
				AddComponent(renderable, &Renderable{
					Image: skellyImg,
				}).
				AddComponent(position, &Position{
					X: mX,
					Y: mY,
					Z: skellyImg.Position().Z,
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
