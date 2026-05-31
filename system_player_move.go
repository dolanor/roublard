package main

import (
	"log/slog"
)

func TryMovePlayer(g *Game) {
	players := g.WorldTags["players"]

	x := g.Extras.currentX
	y := g.Extras.currentY

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y = 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x = 1
	}

	level := g.Map.CurrentLevel

	for _, result := range g.World.Query(players) {
		pos, ok := result.Components[position].(*Position)
		if !ok {
			slog.Error("bad pos", "pos", pos)
			panic("bad pos")
		}

		index := level.GetIndexFromXY(pos.X+x, pos.Y+y)
		//slog.Info("pos", "X", pos.X, "Y", pos.Y, "block", tile.Blocked)

		tile := &level.Tiles[index]
		if tile.Blocked != true {
			pos.X += x
			pos.Y += y
		}

		level.mu.Lock()
		level.PlayerVisible.Compute(level, pos.X, pos.Y, 8)
		level.mu.Unlock()

		updateMapVisibility(level)
	}

	if x != 0 || y != 0 {
		g.Extras.currentX, g.Extras.currentY = 0, 0
		g.Turn = GetNextState(g.Turn)
		g.TurnCounter = 0
	}
}
