package main

import (
	"log/slog"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/gltf"
	"github.com/g3n/engine/math32"
)

func loadElfMesh() core.INode {
	mesh := loadMesh("assets/elf-wizard.glb", 0, 0.01, 0.7+tileHeight)
	meshNode := mesh.GetNode()

	// add torch light
	pointLight := light.NewPoint(&math32.Color{1, .5, 0}, 30)
	pointLight.SetPosition(1, 1, 2)

	meshNode.Add(pointLight)

	return mesh
}

func loadSkeletonMesh() core.INode {
	return loadMesh("assets/skeleton-axe.glb", 0, 0.03, 0.05+tileHeight)
}

func loadGoblinJanitorMesh() core.INode {
	return loadMesh("assets/goblin-janitor.glb", 1, 0.03, -0.8)
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
	meshNode.SetPosition(1, 0.7, zOffset)
	// TODO add to scene somehow
	log.Info("scale", "file_path", path, "scale", mesh.Scale())

	return mesh
}
