package main

import (
	"github.com/dolanor/roublard/assets"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
)

type Tile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Mesh    *graphic.Mesh
}

type Level struct {
	Tiles []Tile
	mm    *assets.MaterialManager
}

func NewLevel() Level {
	mm := assets.NewMaterialManager()
	l := Level{
		mm: mm,
	}

	tiles := l.CreateTiles()
	l.Tiles = tiles

	return l
}

func (l *Level) GetIndexFromXY(x, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

func (l *Level) CreateTiles() []Tile {
	gd := NewGameData()
	tiles := make([]Tile, 0, gd.ScreenHeight*gd.ScreenWidth)

	for x := range gd.ScreenWidth {
		for y := range gd.ScreenHeight {

			if x == 0 || x == gd.ScreenWidth-1 ||
				y == 0 || y == gd.ScreenHeight-1 {
				wall := NewWallTile(x, y, l.mm)
				tile := Tile{
					PixelX:  x,
					PixelY:  y,
					Blocked: true,
					Mesh:    wall,
				}
				tiles = append(tiles, tile)
			} else {
				floor := NewFloorTile(x, y, l.mm)
				floor.SetPosition(float32(x), 0, float32(y))
				tile := Tile{
					PixelX:  x,
					PixelY:  y,
					Blocked: false,
					Mesh:    floor,
				}
				tiles = append(tiles, tile)
			}
		}
	}
	return tiles
}

const tileSideLength = 1
const tileHeight = 0.1

func NewWallTile(x, y int, mm *assets.MaterialManager) *graphic.Mesh {
	height := float32(3)
	geom := geometry.NewBox(tileSideLength, height, tileSideLength)

	mat := mm.Get(assets.MaterialID("wall"))
	mesh := graphic.NewMesh(geom, mat)

	mesh.SetPosition(float32(x), float32(height/2), float32(y))
	return mesh
}

func NewFloorTile(x, y int, mm *assets.MaterialManager) *graphic.Mesh {
	geom := geometry.NewBox(tileSideLength, tileHeight, tileSideLength)

	mesh := graphic.NewMesh(geom, nil)

	mat := mm.Get(assets.MaterialID("floor"))
	mesh.AddGroupMaterial(mat, 2)

	mesh.SetPosition(float32(x), 0, float32(y))
	return mesh
}
