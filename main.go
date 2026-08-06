package main

import (
	"log/slog"
)

func main() {
	log := slog.Default()
	// TODO:
	// maybe we should have the tile meshes separately like InitializeWorld
	// or based on the game world, we generated/reference the meshes
	gameMap := NewGameMap()
	world, tags, meshes := InitializeWorld(gameMap.CurrentLevel)

	game := NewGame(log.With("component", "game"), gameMap, world, tags)

	app, err := NewG3NApp(log.With("component", "g3n"), game, meshes)
	if err != nil {
		panic(err)
	}

	app.SetupKeyboardEventHandlers()
	app.SetupUI()

	app.Run()
}
