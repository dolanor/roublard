package main

import (
	_ "math" // useless, just to make the diff more palatable with rrogue
)

func ProcessRenderables(g *Game, level *Level /* TODO , ??? */) {
	for _, result := range g.World.Query(g.WorldTags["renderables"]) {
		pos := result.Components[position].(*Position)
		node := result.Components[renderable].(*Renderable)

		node.Image.GetNode().SetPosition(float32(pos.X), float32(pos.Z), float32(pos.Y))
		level.mu.Lock()
		if level.PlayerVisible.IsVisible(pos.X, pos.Y) {
			node.Image.SetVisible(true)
		} else {
			node.Image.SetVisible(false)
		}
		level.mu.Unlock()

	}
}
