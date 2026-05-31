package main

import (
	_ "math" // useless, just to make the diff more palatable with rrogue
)

func ProcessRenderables(g *Game, level *Level /* TODO , ??? */) {
	for _, result := range g.World.Query(g.WorldTags["renderables"]) {
		pos := result.Components[position].(*Position)
		node := result.Components[renderable].(*Renderable)

		node.node.GetNode().SetPosition(float32(pos.X), 0.7+tileHeight, float32(pos.Y))
		level.mu.Lock()
		if level.PlayerVisible.IsVisible(pos.X, pos.Y) {
			node.node.SetVisible(true)
		} else {
			node.node.SetVisible(false)
		}
		level.mu.Unlock()

	}
}
