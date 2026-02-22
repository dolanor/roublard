package main

import (
	"log/slog"

	"github.com/g3n/engine/window"
)

type Move string

const (
	MoveLeft  Move = "move_left"
	MoveRight Move = "move_right"
	MoveUp    Move = "move_up"
	MoveDown  Move = "move_down"
	MoveJump  Move = "move_jump"
)

func (g *Game) onKey(evname string, ev any) {

	x, y, z := 0, 0, 0
	kev := ev.(*window.KeyEvent)
	g.log.Info("key pressed", "key", kev.Key)
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
	case window.KeySpace:
		z = 1
	}
	slog.Info("pos", "x", x, "y", y, "z", z)

	level := g.gameMap.CurrentLevel

	for _, res := range g.World.Query(g.WorldTags["players"]) {
		pos, ok := res.Components[position].(*Position)
		if !ok {
			slog.Error("bad pos", "pos", pos)
			panic("bad pos")
		}

		index := level.GetIndexFromXY(pos.X+x, pos.Y+y)
		tile := level.Tiles[index]

		slog.Info("pos", "X", pos.X, "Y", pos.Y, "Z", pos.Z, "block", tile.Blocked)

		if tile.Blocked {
			continue
		}
		pos.X += x
		pos.Y += y
		pos.Z += z
	}
}
