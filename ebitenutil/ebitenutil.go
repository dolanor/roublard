package ebitenutil

import (
	"github.com/dolanor/roublard/assets"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
)

var MaterialManager *assets.MaterialManager

func NewImageFromFile(imgPath string) (*graphic.Mesh, string, error) {
	if MaterialManager == nil {
		MaterialManager = assets.NewMaterialManager()
	}

	var mesh *graphic.Mesh

	switch imgPath {
	case "assets/floor.png":
		mesh = NewFloorMesh(MaterialManager)
	case "assets/wall.png":
		mesh = NewWallMesh(MaterialManager)
	}

	return mesh, "", nil
}

// down is copy from level_g3n to break cyclic dep

const tileSideLength = 1
const tileHeight = 0.1

func NewWallMesh(mm *assets.MaterialManager) *graphic.Mesh {
	height := float32(3)
	geom := geometry.NewBox(tileSideLength, height, tileSideLength)

	mat := mm.Get(assets.MaterialID("wall"))
	mesh := graphic.NewMesh(geom, mat)

	mesh.SetPosition(float32(0), float32(height/2), float32(0))
	mesh.SetVisible(false)
	return mesh
}

func NewFloorMesh(mm *assets.MaterialManager) *graphic.Mesh {
	geom := geometry.NewBox(tileSideLength, tileHeight, tileSideLength)

	mesh := graphic.NewMesh(geom, nil)

	mat := mm.Get(assets.MaterialID("floor"))
	mesh.AddGroupMaterial(mat, 2)

	mesh.SetVisible(false)
	return mesh
}
