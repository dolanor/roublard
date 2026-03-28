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
		g.TryMovePlayers()
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

	for _, res := range g.World.Query(g.WorldTags["players"]) {
		pos, ok := res.Components[position].(*Position)
		if !ok {
			slog.Error("bad pos", "pos", pos)
			panic("bad pos")
		}

		index := level.GetIndexFromXY(pos.X+x, pos.Y+y)
		tile := level.Tiles[index]

		//slog.Info("pos", "X", pos.X, "Y", pos.Y, "block", tile.Blocked)

		if tile.Blocked {
			continue
		}
		pos.X += x
		pos.Y += y
	}
}
