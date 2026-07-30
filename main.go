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

// Update is called each tic.
func (g3nApp *G3NApp) UpdateLogic() error {
	g3nApp.game.TurnCounter++
	if g3nApp.game.Turn == PlayerTurn && g3nApp.game.TurnCounter > 20 {
		TakePlayerAction(g3nApp.game)
	}
	if g3nApp.game.Turn == MonsterTurn {
		UpdateMonster(g3nApp.game)
	}

	ProcessUserLogG3N(g3nApp.game)
	ProcessHUDG3N(g3nApp.game)

	return nil

}

func main() {
	log := slog.Default()
	// TODO:
	// maybe we should have the tile meshes separately like InitializeWorld
	// or based on the game world, we generated/reference the meshes
	gameMap := NewGameMap()
	world, tags, meshes := InitializeWorld(gameMap.CurrentLevel)

	game := NewGame(log.With("component", "game"), gameMap, world, tags)

	app := NewG3NApp(log.With("component", "g3n"), game, meshes)

	app.SetupKeyboardEventHandlers()
	app.SetupUI()

	app.Run()
}
