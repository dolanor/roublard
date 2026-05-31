package main

import "github.com/g3n/engine/core"

type GameMap struct {
	Dungeons     []Dungeon
	CurrentLevel *Level
}

func NewGameMap(scene *core.Node) *GameMap {
	l := NewLevel()
	levels := make([]*Level, 0)
	levels = append(levels, &l)
	d := Dungeon{Name: "default", Levels: levels}
	dungeons := make([]Dungeon, 0)

	for _, t := range l.Tiles {
		scene.Add(t.Image)
	}

	dungeons = append(dungeons, d)
	gm := &GameMap{Dungeons: dungeons, CurrentLevel: &l}
	return gm

}
