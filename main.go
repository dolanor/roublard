package main

import (
	"github.com/bytearena/ecs"
	"github.com/g3n/engine/window"
)

type Game struct {
	Map         *GameMap
	World       *ecs.Manager
	WorldTags   map[string]ecs.Tag
	Turn        TurnState
	TurnCounter int
	Extras      *GameExtras
}

// NewGame creates a new Game Object and initializes the data
// This is a pretty solid refactor candidate for later
func NewGame() *Game {
	extras := NewG3NExtras()
	g := &Game{}
	g.Map = NewGameMap(extras.scene)
	world, tags := InitializeWorld(g.Map.CurrentLevel, extras.scene)

	g.WorldTags = tags
	g.World = world
	g.Turn = PlayerTurn
	g.TurnCounter = 0
	g.Extras = extras
	go g.UpdateLogicLoop()
	return g

}

// Update is called each tic.
func (g *Game) UpdateLogic() error {
	g.TurnCounter++
	if g.Turn == PlayerTurn && g.TurnCounter > 20 {
		TakePlayerAction(g)
	}
	if g.Turn == MonsterTurn {
		UpdateMonster(g)
	}

	ProcessUserLogG3N(g)
	ProcessHUDG3N(g)

	return nil

}

func main() {

	g := NewGame()

	g.Extras.app.Subscribe(window.OnKeyDown, g.onKey)

	width, height := g.Extras.app.GetSize()
	panel := CreatePanel(g.Extras.scene, width, height)
	CreateLogWindow(panel, width/2, height/4)
	CreateHUDWindow(panel, width/2, height/4)
	g.Extras.app.Run(g.Update)
}
