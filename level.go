package main

import (
	"fmt"
	"log/slog"

	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"

	"github.com/dolanor/roublard/assets"
)

type Tile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Mesh    *graphic.Mesh
}

type Level struct {
	Tiles []Tile
	Rooms []Rect

	mm *assets.MaterialManager
}

func NewLevel() Level {
	mm := assets.NewMaterialManager()
	l := Level{
		mm: mm,
	}

	l.generateLevelTiles()

	slog.Info("rooms", "rooms", l.Rooms)
	return l
}

func (l *Level) GetIndexFromXY(x, y int) int {
	// FIXME: reuse another instance of game data
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

func (l *Level) CreateTiles() []Tile {
	gd := NewGameData()
	tiles := make([]Tile, gd.ScreenHeight*gd.ScreenWidth)

	index := 0
	for y := range gd.ScreenHeight {
		for x := range gd.ScreenWidth {
			index = l.GetIndexFromXY(x, y)

			wall := NewWallMesh(x, y, l.mm)
			tile := Tile{
				PixelX:  x,
				PixelY:  y,
				Blocked: true,
				Mesh:    wall,
			}
			tiles[index] = tile
		}
	}
	debugPrintTiles(tiles, gd)
	return tiles
}

func (l *Level) createRoom(room Rect) {
	slog.Info("carving room", "room", room)
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := l.GetIndexFromXY(x, y)

			l.Tiles[index].Blocked = false

			floor := NewFloorMesh(x, y, l.mm)
			l.Tiles[index].Mesh = floor
		}
	}
}

func (l *Level) generateLevelTiles() {
	const (
		minSize  = 6
		maxSize  = 10
		maxRooms = 30
	)

	// TODO: use the same reference to game data
	gd := NewGameData()
	tiles := l.CreateTiles()

	l.Tiles = tiles

	for i := 0; i < maxRooms; i++ {
		w := GetRandomBetween(minSize, maxSize)
		h := GetRandomBetween(minSize, maxSize)
		x := GetDiceRoll(gd.ScreenWidth-w-1) - 1
		y := GetDiceRoll(gd.ScreenHeight-w-1) - 1

		newRoom := NewRect(x, y, w, h)

		slog.Info("creating room", "i", i, "room", newRoom)
		okToAdd := true

		for _, otherRoom := range l.Rooms {
			if newRoom.Intersect(otherRoom) {
				okToAdd = false
				break
			}
		}

		if okToAdd {
			l.createRoom(newRoom)
			l.Rooms = append(l.Rooms, newRoom)
		}
	}

}

func debugPrintTiles(tiles []Tile, gameData GameData) {

	fmt.Println("===============================")

	for i, t := range tiles {
		if i%gameData.ScreenWidth == 0 {
			fmt.Println()
		}
		tileChar := "."
		if t.Blocked {
			tileChar = "#"
		}
		fmt.Printf("%s", tileChar)

	}

	fmt.Println("\n\n===============================")
}

const tileSideLength = 1
const tileHeight = 0.1

func NewWallMesh(x, y int, mm *assets.MaterialManager) *graphic.Mesh {
	height := float32(3)
	geom := geometry.NewBox(tileSideLength, height, tileSideLength)

	mat := mm.Get(assets.MaterialID("wall"))
	mesh := graphic.NewMesh(geom, mat)

	mesh.SetPosition(float32(x), float32(height/2), float32(y))
	return mesh
}

func NewFloorMesh(x, y int, mm *assets.MaterialManager) *graphic.Mesh {
	geom := geometry.NewBox(tileSideLength, tileHeight, tileSideLength)

	mesh := graphic.NewMesh(geom, nil)

	mat := mm.Get(assets.MaterialID("floor"))
	mesh.AddGroupMaterial(mat, 2)

	mesh.SetPosition(float32(x), 0, float32(y))
	return mesh
}
