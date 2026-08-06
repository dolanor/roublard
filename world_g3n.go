package main

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/gltf"
	"github.com/g3n/engine/math32"
)

const (
	ModelFilePathElfWizard     = "assets/elf-wizard.glb"
	ModelFilePathSkeleton      = "assets/skeleton-axe.glb"
	ModelFilePathGoblinJanitor = "assets/goblin-janitor.glb"
)

func loadElfMesh() core.INode {
	mesh := loadMesh(ModelFilePathElfWizard, 0, 0.01, 0.7+tileHeight)

	// add torch light
	pointLight := light.NewPoint(&math32.Color{R: 1, G: .5}, 30)
	pointLight.SetPosition(1, 1, 2)

	meshNode := mesh.GetNode()
	meshNode.Add(pointLight)

	return mesh
}

func loadSkeletonMesh() core.INode {
	return loadMesh(ModelFilePathSkeleton, 0, 0.03, 0.05+tileHeight)
}

func loadGoblinJanitorMesh() core.INode {
	return loadMesh(ModelFilePathGoblinJanitor, 1, 0.03, 0)
}

type MeshCache struct {
	mu     sync.Mutex
	meshes map[string]*graphic.Mesh
}

var meshCache = MeshCache{meshes: map[string]*graphic.Mesh{}}

func loadMesh(path string, meshIndex int, scaleFactor float32, zOffset float32) core.INode {
	// FIXME: use the game logger
	log := slog.Default()

	mesh, ok := meshCache.meshes[path]
	if !ok {
		slog.Warn("mesh cache miss", "path", path)
		model, err := gltf.ParseBin(path)
		if err != nil {
			panic(err)
		}
		log.Info("load model", "len(meshes)", len(model.Meshes))

		inode, err := model.LoadMesh(meshIndex)
		if err != nil {
			panic(err)
		}

		m, ok := inode.(*graphic.Mesh)
		if !ok {
			slog.Error("the mesh is not a *graphic.Mesh", "type", fmt.Sprintf("%T", inode))
			panic("bad mesh")
		}

		mesh = m

		meshCache.meshes[path] = mesh
	} else {
		slog.Info("mesh cache hit", "path", path)
	}

	// FIXME: find out why the player mesh cloned will mess up the lighting.
	if path != ModelFilePathElfWizard {
		mesh = CloneAndPosition(mesh, 1, 1)
	}

	meshNode := mesh.GetNode()
	meshNode.SetScale(scaleFactor, scaleFactor, scaleFactor)
	// depends on the model size I suppose
	meshNode.SetPosition(1, 0.7, zOffset)
	// TODO add to scene somehow
	log.Info("scale", "file_path", path, "scale", mesh.Scale())

	return mesh
}
