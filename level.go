package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"

	"github.com/dolanor/roublard/assets"
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
	height := float32(3)
	geom := geometry.NewBox(tileSideLength, height, tileSideLength)
	color := math32.NewColor("White")
	mat := material.NewStandard(color)
	// FIXME: remove once camera is debugged
	//color.R = float32(x) / 80
	//color.G = float32(y) / 50

	//mat := material.NewPhysical()
	//var eg errgroup.Group
	//eg.Go(func() error { mat.SetEmissiveFactor(math32.NewColor("white")); return nil })
	//eg.Go(func() error { mat.SetBaseColorMap(textures.WoodInlaidDiffuse()); return nil })
	//eg.Go(func() error { mat.SetMetallicRoughnessMap(textures.WoodInlaidRough()); return nil })
	//eg.Go(func() error { mat.SetNormalMap(textures.WoodInlaidNormal()); return nil })

	//err := eg.Wait()
	//if err != nil {
	//	panic(err)
	//}

	tex := assets.Wall()
	mat.AddTexture(tex)
	mesh := graphic.NewMesh(geom, mat)

	mesh.SetPosition(float32(x), float32(height/2), float32(y))
	return mesh
}

func NewFloorTile(x, y int) *graphic.Mesh {
	geom := geometry.NewBox(tileSideLength, .1, tileSideLength)
	color := math32.NewColor("White")
	// FIXME: remove once camera is debugged
	//color.R = float32(x) / 80
	//color.G = float32(y) / 50

	mat := material.NewStandard(color)
	tex := assets.Floor()
	mat.AddTexture(tex)
	mesh := graphic.NewMesh(geom, mat)

	mesh.SetPosition(float32(x), 0, float32(y))
	return mesh
}
