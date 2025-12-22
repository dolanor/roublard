package main

import (
	"log/slog"

	"github.com/bytearena/ecs"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/gltf"
	"github.com/g3n/engine/math32"
)

var position *ecs.Component
var renderable *ecs.Component

func InitWorld(scene *core.Node) (*ecs.Manager, map[string]ecs.Tag) {
	tags := map[string]ecs.Tag{}

	mgr := ecs.NewManager()

	// WARNING: this is global state
	position = mgr.NewComponent()
	renderable = mgr.NewComponent()

	player := mgr.NewComponent()
	movable := mgr.NewComponent()

	mesh := loadElfMesh()
	scene.Add(mesh)

	// Define the elf wizard in the ECS
	mgr.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{
			node: mesh,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: 40,
			Y: 25,
		})

	pointLight := light.NewPoint(&math32.Color{1, .5, 0}, 30)
	pointLight.SetPosition(1, 1, 2)
	scene.Add(pointLight)

	// Add the movable light (invisible torch for now) in the ECS
	mgr.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{
			node: pointLight,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: 40,
			Y: 25,
		})

	players := ecs.BuildTag(player, position)
	tags["players"] = players

	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables

	return mgr, tags
}

func loadElfMesh() core.INode {
	// FIXME: use the game logger
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
