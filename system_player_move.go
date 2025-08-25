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
)

func (g *Game) onKey(evname string, ev any) {
	for _, res := range g.World.Query(g.WorldTags["players"]) {
		pos, ok := res.Components[position].(*Position)
		if !ok {
			slog.Error("bad pos", "pos", pos)
			panic("bad pos")
		}

		kev := ev.(*window.KeyEvent)
		g.log.Info("key pressed", "key", kev.Key)
		switch kev.Key {
		case window.KeyE:
			pos.Y--
		case window.KeyD:
			pos.Y++
		case window.KeyS:
			pos.X--
		case window.KeyF:
			pos.X++
		}
		slog.Info("pos", "x", pos.X, "y", pos.Y)
	}
}
