package main

import (
	"fmt"
	"log/slog"

	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/norendren/go-fov/fov"

	"github.com/dolanor/roublard/assets"
)

type Tile struct {
	PixelX     int
	PixelY     int
	Blocked    bool
	Mesh       *graphic.Mesh
	IsRevealed bool
	IsWall     bool
}

type Level struct {
	Tiles []Tile
	Rooms []Rect

	mm            *assets.MaterialManager
	gameData      GameData
	PlayerVisible *fov.View
}

func NewLevel() Level {
	mm := assets.NewMaterialManager()
	l := Level{
		mm:            mm,
		gameData:      NewGameData(),
		PlayerVisible: fov.New(),
	}

	l.generateLevelTiles()

	slog.Info("rooms", "rooms", l.Rooms)
	return l
}

func (l *Level) GetIndexFromXY(x, y int) int {
	gd := l.gameData
	return (y * gd.ScreenWidth) + x
}

func (l *Level) CreateTiles() []Tile {
	gd := l.gameData
	tiles := make([]Tile, gd.ScreenHeight*gd.ScreenWidth)

	index := 0
	for y := range gd.ScreenHeight {
		for x := range gd.ScreenWidth {
			index = l.GetIndexFromXY(x, y)

			wall := NewWallMesh(x, y, l.mm)
			wall.SetVisible(false)
			tile := Tile{
				PixelX:  x,
				PixelY:  y,
				Blocked: true,
				Mesh:    wall,
				IsWall:  true,
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
			l.Tiles[index].IsWall = false

			floor := NewFloorMesh(x, y, l.mm)
			floor.SetVisible(false)
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
	containsRoom := false

	gd := l.gameData
	tiles := l.CreateTiles()

	l.Tiles = tiles

	for i := 0; i < maxRooms; i++ {
		w := GetRandomBetween(minSize, maxSize)
		h := GetRandomBetween(minSize, maxSize)
		x := GetDiceRoll(gd.ScreenWidth - w - 1)
		y := GetDiceRoll(gd.ScreenHeight - w - 1)

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
			if containsRoom {
				newX, newY := newRoom.Center()
				prevX, prevY := l.Rooms[len(l.Rooms)-1].Center()

				coinFlip := GetDiceRoll(2)
				if coinFlip == 2 {
					l.createHorizontalTunnel(prevX, newX, prevY)
					l.createVerticalTunnel(prevY, newY, newX)
				} else {
					l.createHorizontalTunnel(prevX, newX, newY)
					l.createVerticalTunnel(prevY, newY, prevX)
				}
			}

			l.Rooms = append(l.Rooms, newRoom)
			containsRoom = true
		}
	}

}

func (l *Level) createHorizontalTunnel(x1, x2, y int) {
	gd := l.gameData
	for x := min(x1, x2); x < max(x1, x2)+1; x++ {
		index := l.GetIndexFromXY(x, y)

		if index > 0 && index < gd.ScreenWidth*gd.ScreenHeight {
			l.Tiles[index].Blocked = false
			l.Tiles[index].IsWall = false
			floor := NewFloorMesh(x, y, l.mm)
			floor.SetVisible(false)

			l.Tiles[index].Mesh = floor
		}
	}
}

func (l *Level) createVerticalTunnel(y1, y2, x int) {
	gd := l.gameData
	for y := min(y1, y2); y < max(y1, y2)+1; y++ {
		index := l.GetIndexFromXY(x, y)

		if index > 0 && index < gd.ScreenWidth*gd.ScreenHeight {
			l.Tiles[index].Blocked = false
			l.Tiles[index].IsWall = false
			floor := NewFloorMesh(x, y, l.mm)
			floor.SetVisible(false)

			l.Tiles[index].Mesh = floor
		}
	}
}

func (level Level) InBounds(x, y int) bool {
	gd := NewGameData()
	if x < 0 || x > gd.ScreenWidth || y < 0 || y > gd.ScreenHeight {
		return false
	}
	return true
}

// TODO: Change this to check for WALL, not blocked
func (level Level) IsOpaque(x, y int) bool {
	idx := level.GetIndexFromXY(x, y)
	return level.Tiles[idx].Blocked
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
