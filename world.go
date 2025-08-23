package main

import (
	"github.com/bytearena/ecs"
)

func InitWorld() (*ecs.Manager, map[string]ecs.Tag) {
	tags := map[string]ecs.Tag{}

	mgr := ecs.NewManager()

	player := mgr.NewComponent()
	position := mgr.NewComponent()
	renderable := mgr.NewComponent()
	movable := mgr.NewComponent()

	mgr.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, Renderable{}).
		AddComponent(movable, Movable{}).
		AddComponent(position, Position{
			X: 40,
			Y: 25,
		})

	players := ecs.BuildTag(player, position)
	tags["players"] = players

	return mgr, tags
}
