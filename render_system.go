package main

import (
	"log/slog"
	"math"
	"time"
)

func ProcessRenderables(g *Game, level *Level) {
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

func animateDeath(renderable *Renderable) {
	maxAnimTime := time.Now().Add(1 * time.Second)
	var i int
	for t := range time.Tick(10 * time.Millisecond) {
		i++
		if t.After(maxAnimTime) {
			slog.Info("passed max anim time")
			break
		}
		renderable.Image.GetNode().SetRotationY(float32(2*i) * math.Pi * 2 / 100)
	}
	renderable.Image.GetNode().SetRotationX(math.Pi / 2)
}
