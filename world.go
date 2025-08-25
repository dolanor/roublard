package main

import (
	"log/slog"

	"github.com/bytearena/ecs"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/loader/gltf"
)

func InitWorld(scene *core.Node) (*ecs.Manager, map[string]ecs.Tag) {
	tags := map[string]ecs.Tag{}

	mgr := ecs.NewManager()

	player := mgr.NewComponent()
	position := mgr.NewComponent()
	renderable := mgr.NewComponent()
	movable := mgr.NewComponent()

	node := loadElf()

	scene.Add(node)

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

func loadElf() core.INode {
	log := slog.Default()

	model, err := gltf.ParseBin("assets/elf-wizard.glb")
	if err != nil {
		panic(err)
	}
	log.Info("load model", "len(meshes)", len(model.Meshes))

	mesh, err := model.LoadMesh(0)
	if err != nil {
		panic(err)
	}

	meshNode := mesh.GetNode()
	meshNode.SetScale(0.01, 0.01, 0.01)
	// depends on the model size I suppose
	meshNode.SetPosition(1, 0.7+tileHeight, 1)
	// TODO add to scene somehow
	log.Info("elf wizard", "scale", mesh.Scale())

	return mesh
}
