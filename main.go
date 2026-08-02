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

	app := NewG3NApp(log.With("component", "g3n"), game, meshes)

	app.SetupKeyboardEventHandlers()
	app.SetupUI()

	app.Run()
}
