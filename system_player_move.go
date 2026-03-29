package main

import (
	"log/slog"
	"time"

	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/window"
)

type Move string

const (
	MoveLeft  Move = "move_left"
	MoveRight Move = "move_right"
	MoveUp    Move = "move_up"
	MoveDown  Move = "move_down"
)

func (g *Game) logicUpdateLoop() {
	for range time.Tick(time.Second / 60) {
		g.TurnCounter++
		if g.Turn == PlayerTurn && g.TurnCounter > 20 {
			g.TryMovePlayers()
		}
		// Obviously just for now
		g.Turn = PlayerTurn
	}
}

func (g *Game) onKey(evname string, ev any) {
	g.currentX, g.currentY = g.processKeys(ev)
}

func (g *Game) processKeys(ev any) (x, y int) {

	kev := ev.(*window.KeyEvent)
	switch kev.Key {
	case window.KeyE:
		y = -1
	case window.KeyD:
		y = 1
	case window.KeyS:
		x = -1
	case window.KeyF:
		x = 1

	case window.KeyM:
		if kev.Mods == window.ModControl {
			g.app.Exit()
		}

	case window.KeyU:
		if kev.Mods == window.ModControl {
			if g.orthoToggle {
				g.cam.SetProjection(camera.Orthographic)
				g.orthoToggle = !g.orthoToggle
			} else {
				g.cam.SetProjection(camera.Perspective)
				g.orthoToggle = !g.orthoToggle
			}

		}
	}

	return x, y
}

func (g *Game) TryMovePlayers() {
	level := g.gameMap.CurrentLevel

	x, y := g.currentX, g.currentY
	gd := level.gameData

	for _, res := range g.World.Query(g.WorldTags["players"]) {
		pos, ok := res.Components[position].(*Position)
		if !ok {
			slog.Error("bad pos", "pos", pos)
			panic("bad pos")
		}

		index := level.GetIndexFromXY(pos.X+x, pos.Y+y)
		tile := &level.Tiles[index]

		//slog.Info("pos", "X", pos.X, "Y", pos.Y, "block", tile.Blocked)

		if tile.Blocked {
			continue
		}
		pos.X += x
		pos.Y += y
		level.PlayerVisible.Compute(level, pos.X, pos.Y, 8)

		// We decide to check for every tile in the level if it should be rendered or not
		for x := 0; x < gd.ScreenWidth; x++ {
			for y := 0; y < gd.ScreenHeight; y++ {
				index := level.GetIndexFromXY(x, y)
				tile := &level.Tiles[index]

				if level.PlayerVisible.IsVisible(x, y) {
					tile.Mesh.SetRenderable(true)
				} else {
					tile.Mesh.SetRenderable(false)
				}
			}
		}
	}

	if x != 0 || y != 0 {
		g.currentX, g.currentY = 0, 0
		g.Turn = GetNextState(g.Turn)
		g.TurnCounter = 0
	}
}
