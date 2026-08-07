package main

import (
	"github.com/g3n/engine/graphic"

	"codeberg.org/dolanor/roublard/assets"
)

func NewTileMeshFromFile(mm *assets.MaterialManager, imgPath string) (*graphic.Mesh, error) {
	var mesh *graphic.Mesh

	switch imgPath {
	case "assets/floor.png":
		mesh = NewFloorMesh(mm)
	case "assets/wall.png":
		mesh = NewWallMesh(mm)
	}

	return mesh, nil
}
