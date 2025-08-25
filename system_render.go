package main

import (
	"log/slog"
)

func ProcessRenderables(game *Game, level Level /* TODO , ??? */) {
	for _, res := range game.World.Query(game.WorldTags["renderables"]) {
		_ = res
		pos, ok := res.Components[position].(*Position)
		if !ok {
			slog.Error("bad pos", "pos", pos)
		}
		node, ok := res.Components[renderable].(*Renderable)
		if !ok {
			slog.Error("bad node", "node", node)
		}

		node.node.GetNode().SetPosition(float32(pos.X), 0.7+tileHeight, float32(pos.Y))
	}
}
