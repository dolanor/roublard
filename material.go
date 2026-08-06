package main

import (
	"codeberg.org/dolanor/roublard/assets"
	"github.com/g3n/engine/graphic"
)

func NewMeshFromFile(mm *assets.MaterialManager, imgPath string) (*graphic.Mesh, error) {

	var mesh *graphic.Mesh

	switch imgPath {
	case "assets/floor.png":
		mesh = NewFloorMesh(mm)
	case "assets/wall.png":
		mesh = NewWallMesh(mm)
	}

	return mesh, nil
}
