package main

import (
	"log/slog"
	"sync"

	"github.com/g3n/engine/graphic"
	"github.com/norendren/go-fov/fov"

	"github.com/dolanor/roublard/assets"
)

type TileType int

const (
	WALL TileType = iota
	FLOOR
)

// Level holds the tile information for a complete dungeon level.
type Level struct {
	Tiles         []*MapTile
	Rooms         []Rect
	PlayerVisible *fov.View
	mu            sync.Mutex

	mm       *assets.MaterialManager
	gameData GameData
}

// MapTile is a single Tile on a given level
type MapTile struct {
	PixelX     int
	PixelY     int
	Blocked    bool
	Image      *graphic.Mesh
	IsRevealed bool
	IsWall     bool
	TileType   TileType
}

func NewLevel() Level {
	l := Level{}
	l.mm = assets.NewMaterialManager()

	rooms := make([]Rect, 0)
	l.Rooms = rooms
	l.GenerateLevelTiles()
	l.PlayerVisible = fov.New()
	slog.Info("rooms", "rooms", l.Rooms)
	// TODO: reuse this in the future
	l.gameData = NewGameData()
	return l
}

// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func (level *Level) GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

// GenerateLevelTiles creates a new Dungeon Level Map.
func (level *Level) GenerateLevelTiles() {
	MIN_SIZE := 6
	MAX_SIZE := 10
	MAX_ROOMS := 30

	gd := NewGameData()
	tiles := level.createTiles()
	level.Tiles = tiles
	contains_rooms := false

	for idx := 0; idx < MAX_ROOMS; idx++ {
		w := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		h := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		x := GetDiceRoll(gd.ScreenWidth - w - 1)
		y := GetDiceRoll(gd.ScreenHeight - h - 1)
		new_room := NewRect(x, y, w, h)

		slog.Info("creating room", "i", idx, "room", new_room)
		okToAdd := true

		for _, otherRoom := range level.Rooms {
			if new_room.Intersect(otherRoom) {
				okToAdd = false
				break
			}
		}
		if okToAdd {
			level.createRoom(new_room)
			if contains_rooms {
				newX, newY := new_room.Center()
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

			level.Rooms = append(level.Rooms, new_room)
			contains_rooms = true
		}
	}

}

func (level *Level) createHorizontalTunnel(x1 int, x2 int, y int) {
	gd := NewGameData()
	for x := min(x1, x2); x < max(x1, x2)+1; x++ {
		index := level.GetIndexFromXY(x, y)
		if index > 0 && index < gd.ScreenWidth*gd.ScreenHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].IsWall = false
			level.Tiles[index].TileType = FLOOR
			floor := NewFloorMesh(x, y, level.mm)
			level.Tiles[index].Image = floor

		}
	}
}

func (level *Level) createVerticalTunnel(y1 int, y2 int, x int) {
	gd := NewGameData()
	for y := min(y1, y2); y < max(y1, y2)+1; y++ {
		index := level.GetIndexFromXY(x, y)
		if index > 0 && index < gd.ScreenWidth*gd.ScreenHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].IsWall = false
			level.Tiles[index].TileType = FLOOR
			floor := NewFloorMesh(x, y, level.mm)
			level.Tiles[index].Image = floor
		}
	}
}

// createTiles creates a map of all walls as a baseline for carving out a level.
func (level *Level) createTiles() []*MapTile {
	gd := NewGameData()
	tiles := make([]*MapTile, gd.ScreenHeight*gd.ScreenWidth)
	index := 0
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			index = level.GetIndexFromXY(x, y)
			wall := NewWallMesh(x, y, level.mm)
			tile := MapTile{
				PixelX:     x,
				PixelY:     y,
				Blocked:    true,
				Image:      wall,
				IsRevealed: false,
				TileType:   WALL,
				IsWall:     true,
			}
			tiles[index] = &tile
		}
	}
	debugPrintTiles(tiles, gd)
	return tiles
}

func (level *Level) createRoom(room Rect) {
	slog.Info("carving room", "room", room)
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)

			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			floor := NewFloorMesh(x, y, level.mm)
			level.Tiles[index].Image = floor
			level.Tiles[index].IsWall = false
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
	return level.Tiles[idx].TileType == WALL
}

// Max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
