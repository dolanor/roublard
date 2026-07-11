package main

import (
	"github.com/bytearena/ecs"
	"github.com/dolanor/roublard/ebiten"
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

	return nil

}

// Draw is called each draw cycle and is where we will blit.
func (g *Game) Draw(screen *ebiten.Image) {
	//Draw the Map
	level := g.Map.CurrentLevel
	level.DrawLevel(screen)
	ProcessRenderables(g, level)
}

// Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) {
	gd := NewGameData()
	return gd.TileWidth * gd.ScreenWidth, gd.TileHeight * gd.ScreenHeight

}

func main() {

	g := NewGame()
	ebiten.SetWindowResizable(true)

	ebiten.SetWindowTitle("Tower")

	g.Extras.app.Subscribe(window.OnKeyDown, g.onKey)

	g.Extras.app.Run(g.Update)
}
