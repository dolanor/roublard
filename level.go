package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}

func NewGameData() GameData {
	return GameData{
		ScreenWidth:  80,
		ScreenHeight: 50,
		TileWidth:    16,
		TileHeight:   16,
	}
}

type Tile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Mesh    *graphic.Mesh
}

func GetIndexFromXY(x, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

func CreateTiles() []Tile {
	gd := NewGameData()
	tiles := []Tile{}

	for x := range gd.ScreenWidth {
		for y := range gd.ScreenHeight {

			if x == 0 || x == gd.ScreenWidth-1 ||
				y == 0 || y == gd.ScreenHeight-1 {
				wall := NewWallTile(x, y)
				tile := Tile{
					PixelX:  x,
					PixelY:  y,
					Blocked: true,
					Mesh:    wall,
				}
				tiles = append(tiles, tile)
			} else {
				floor := NewFloorTile(x, y)
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

func NewWallTile(x, y int) *graphic.Mesh {
	geom := geometry.NewBox(tileSideLength, .1, tileSideLength)
	color := math32.NewColor("DarkBlue")
	// FIXME: use a variable instead of magic number
	color.R = float32(x) / 80
	color.G = float32(y) / 50

	mat := material.NewStandard(color)
	mesh := graphic.NewMesh(geom, mat)

	mesh.SetPosition(float32(x), 0, float32(y))
	return mesh
}

func NewFloorTile(x, y int) *graphic.Mesh {
	geom := geometry.NewBox(tileSideLength, .1, tileSideLength)
	color := math32.NewColor("DarkRed")
	// FIXME: use a variable instead of magic number
	color.R = float32(x) / 80
	color.G = float32(y) / 50

	mat := material.NewStandard(color)
	mesh := graphic.NewMesh(geom, mat)

	mesh.SetPosition(float32(x), 0, float32(y))
	return mesh
}
