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
var monster *ecs.Component

func InitWorld(scene *core.Node, startLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := map[string]ecs.Tag{}

	mgr := ecs.NewManager()

	// WARNING: this is global state
	position = mgr.NewComponent()
	renderable = mgr.NewComponent()

	player := mgr.NewComponent()
	movable := mgr.NewComponent()

	elfMesh := loadElfMesh()
	scene.Add(elfMesh)
	elfMesh.SetVisible(true)

	monster = mgr.NewComponent()

	// Get First Room
	startRoom := startLevel.Rooms[0]
	x, y := startRoom.Center()

	// Define the elf wizard in the ECS
	mgr.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{
			node: elfMesh,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: x,
			Y: y,
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
			X: x,
			Y: y,
		})

	players := ecs.BuildTag(player, position)
	tags["players"] = players

	for _, room := range startLevel.Rooms {
		if room.X1 != startRoom.X1 {
			monsterMesh := loadElfMesh()
			//monsterMesh := loadSkeletonMesh()
			// Make it taller to separate from player
			//monsterMesh.GetNode().SetScale(0.01, 0.02, 0.01)
			//monsterMesh.GetNode().UpdateMatrix()
			monsterMesh.SetVisible(false)

			scene.Add(monsterMesh)

			mX, mY := room.Center()

			mgr.NewEntity().
				AddComponent(monster, Monster{}).
				AddComponent(renderable, &Renderable{
					node: monsterMesh,
				}).
				AddComponent(position, &Position{
					X: mX,
					Y: mY,
				})
		}
	}

	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables

	return mgr, tags
}

func loadElfMesh() core.INode {
	return loadMesh("assets/elf-wizard.glb", 0, 0.01, tileHeight)
}

func loadSkeletonMesh() core.INode {
	return loadMesh("assets/skeleton-axe-wielder.glb", 1, 0.03, -1.8)
}

func loadMesh(path string, meshIndex int, scaleFactor float32, zOffset float32) core.INode {
	// FIXME: use the game logger
	log := slog.Default()

	model, err := gltf.ParseBin(path)
	if err != nil {
		panic(err)
	}
	log.Info("load model", "len(meshes)", len(model.Meshes))

	mesh, err := model.LoadMesh(meshIndex)
	if err != nil {
		panic(err)
	}

	meshNode := mesh.GetNode()
	meshNode.SetScale(scaleFactor, scaleFactor, scaleFactor)
	// depends on the model size I suppose
	meshNode.SetPosition(1, 0.7+zOffset, 1)
	// TODO add to scene somehow
	log.Info("scale", "file_path", path, "scale", mesh.Scale())

	return mesh
}
