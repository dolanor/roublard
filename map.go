package main

// GameMap holds all the level and aggregate information for the entire world.
type GameMap struct {
	Dungeons     []Dungeon
	CurrentLevel *Level
}

// NewGameMap creates a new set of maps for the entire game.
func NewGameMap() *GameMap {
	//Return a new game map of a single level for now
	l := NewLevel()

	d := Dungeon{
		Name:   "default",
		Levels: []*Level{l},
	}

	gm := &GameMap{
		Dungeons:     []Dungeon{d},
		CurrentLevel: l,
	}
	return gm

}
