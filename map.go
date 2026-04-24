package main

import "github.com/g3n/engine/core"

type GameMap struct {
	Dungeons     []Dungeon
	CurrentLevel *Level
}

func NewGameMap(scene *core.Node) GameMap {
	var dungeons []Dungeon
	var levels []*Level
	l := NewLevel()
	levels = append(levels, &l)

	tiles := l.Tiles
	for _, t := range tiles {
		scene.Add(t.Mesh)
	}

	d := Dungeon{
		Name:   "default",
		Levels: levels,
	}

	dungeons = append(dungeons, d)

	return GameMap{
		Dungeons:     dungeons,
		CurrentLevel: &l,
	}
}
