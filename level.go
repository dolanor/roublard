package main

import (
	"log"
	"sync"

	"github.com/g3n/engine/graphic"
	"github.com/norendren/go-fov/fov"

	"codeberg.org/dolanor/roublard/assets"
)

type TileType int

var floor *graphic.Mesh
var wall *graphic.Mesh

var levelHeight int

const (
	TileTypeWall TileType = iota
	TileTypeFloor
)

// Level holds the tile information for a complete dungeon level.
type Level struct {
	Tiles         []*Tile
	Rooms         []Rect
	PlayerVisible *fov.View
	mu            sync.Mutex
}

// Tile is a single Tile on a given level
type Tile struct {
	X       int
	Y       int
	Blocked bool
	// FIXME: move this field into a dedicated g3n tile type
	Mesh       *graphic.Mesh
	IsRevealed bool
	IsWall     bool
	TileType   TileType
}

// NewLevel creates a new game level in a dungeon.
func NewLevel() *Level {

	l := Level{
		Rooms:         []Rect{},
		PlayerVisible: fov.New(),
	}

	l.GenerateLevelTiles()

	return &l
}

func loadTileMeshes(mm *assets.MaterialManager) {
	if floor != nil && wall != nil {
		return
	}

	var err error
	floor, err = NewMeshFromFile(mm, "assets/floor.png")
	if err != nil {
		log.Fatal(err)
	}

	wall, err = NewMeshFromFile(mm, "assets/wall.png")
	if err != nil {
		log.Fatal(err)
	}
}

// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func (level *Level) GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

// GenerateLevelTiles creates a new Dungeon Level Map.
func (level *Level) GenerateLevelTiles() {
	const (
		minSize  = 6
		maxSize  = 10
		maxRooms = 30
	)

	gd := NewGameData()
	levelHeight = gd.ScreenHeight - gd.UIHeight
	tiles := level.createTiles()
	level.Tiles = tiles
	containsRooms := false

	for range maxRooms {
		w := GetRandomBetween(minSize, maxSize)
		h := GetRandomBetween(minSize, maxSize)
		x := GetDiceRoll(gd.ScreenWidth - w - 1)
		y := GetDiceRoll(levelHeight - h - 1)

		newRoom := NewRect(x, y, w, h)

		okToAdd := true
		for _, otherRoom := range level.Rooms {
			if newRoom.Intersect(otherRoom) {
				okToAdd = false
				break
			}
		}

		if okToAdd {
			level.createRoom(newRoom)
			debugPrintTiles(tiles, gd)
			if containsRooms {
				newX, newY := newRoom.Center()
				prevX, prevY := level.Rooms[len(level.Rooms)-1].Center()
				coinflip := GetDiceRoll(2)
				if coinflip == 2 {
					level.createHorizontalTunnel(prevX, newX, prevY)
					level.createVerticalTunnel(prevY, newY, newX)
				} else {
					level.createHorizontalTunnel(prevX, newX, newY)
					level.createVerticalTunnel(prevY, newY, prevX)
				}
			}

			level.Rooms = append(level.Rooms, newRoom)
			containsRooms = true
		}
	}
	debugPrintTiles(tiles, gd)
}

func (level *Level) createHorizontalTunnel(x1 int, x2 int, y int) {
	gd := NewGameData()
	for x := min(x1, x2); x < max(x1, x2)+1; x++ {
		index := level.GetIndexFromXY(x, y)
		if index > 0 && index < gd.ScreenWidth*levelHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].IsWall = false
			level.Tiles[index].TileType = TileTypeFloor
		}
	}
}

func (level *Level) createVerticalTunnel(y1 int, y2 int, x int) {
	gd := NewGameData()
	for y := min(y1, y2); y < max(y1, y2)+1; y++ {
		index := level.GetIndexFromXY(x, y)
		if index > 0 && index < gd.ScreenWidth*levelHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].IsWall = false
			level.Tiles[index].TileType = TileTypeFloor
		}
	}
}

// createTiles creates a map of all walls as a baseline for carving out a level.
func (level *Level) createTiles() []*Tile {
	gd := NewGameData()
	tiles := make([]*Tile, levelHeight*gd.ScreenWidth)
	index := 0
	for x := range gd.ScreenWidth {
		for y := range levelHeight {
			index = level.GetIndexFromXY(x, y)
			tile := Tile{
				X:          x,
				Y:          y,
				Blocked:    true,
				IsRevealed: false,
				TileType:   TileTypeWall,
				IsWall:     true,
			}
			tiles[index] = &tile
		}
	}
	return tiles
}

func (level *Level) createRoom(room Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = TileTypeFloor
			level.Tiles[index].IsWall = false
		}
	}
}

func (*Level) InBounds(x, y int) bool {
	gd := NewGameData()
	if x < 0 || x > gd.ScreenWidth || y < 0 || y > levelHeight {
		return false
	}

	return true
}

func (level *Level) IsOpaque(x, y int) bool {
	idx := level.GetIndexFromXY(x, y)
	return level.Tiles[idx].TileType == TileTypeWall
}
