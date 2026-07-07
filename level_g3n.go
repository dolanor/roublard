package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"

	"github.com/dolanor/roublard/assets"
)

const tileSideLength = 1
const tileHeight = 0.1

func NewWallMesh(x, y int, mm *assets.MaterialManager) *graphic.Mesh {
	height := float32(3)
	geom := geometry.NewBox(tileSideLength, height, tileSideLength)

	mat := mm.Get(assets.MaterialID("wall"))
	mesh := graphic.NewMesh(geom, mat)

	mesh.SetPosition(float32(x), float32(height/2), float32(y))
	mesh.SetVisible(false)
	return mesh
}

func NewFloorMesh(x, y int, mm *assets.MaterialManager) *graphic.Mesh {
	geom := geometry.NewBox(tileSideLength, tileHeight, tileSideLength)

	mesh := graphic.NewMesh(geom, nil)

	mat := mm.Get(assets.MaterialID("floor"))
	mesh.AddGroupMaterial(mat, 2)

	mesh.SetPosition(float32(x), 0, float32(y))
	mesh.SetVisible(false)
	return mesh
}

func CloneAndPosition(mesh *graphic.Mesh, x, y int) *graphic.Mesh {
	m := mesh.Clone().(*graphic.Mesh)
	m.SetPositionX(float32(x))
	m.SetPositionZ(float32(y))

	return m
}
