package main

import (
	"log/slog"

	"github.com/bytearena/ecs"
)

type Game struct {
	Map         *GameMap
	World       *ecs.Manager
	WorldTags   map[string]ecs.Tag
	Turn        TurnState
	TurnCounter int

	// FIXME: maybe protect it with a mutex
	currentX int
	currentY int

	log *slog.Logger
}

func NewGame(log *slog.Logger, gameMap *GameMap, world *ecs.Manager, tags map[string]ecs.Tag) *Game {
	g := &Game{
		log:         log,
		Map:         gameMap,
		WorldTags:   tags,
		World:       world,
		Turn:        PlayerTurn,
		TurnCounter: 0,
	}

	return g
}
