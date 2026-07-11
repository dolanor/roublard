package main

import "github.com/dolanor/roublard/ebiten"

func TakePlayerAction(g *Game) {
	players := g.WorldTags["players"]
	turnTaken := false

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
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		turnTaken = true
	}

	level := g.Map.CurrentLevel

	for _, result := range g.World.Query(players) {
		pos := result.Components[position].(*Position)
		index := level.GetIndexFromXY(pos.X+x, pos.Y+y)

		tile := level.Tiles[index]
		if tile.Blocked != true {
			level.Tiles[level.GetIndexFromXY(pos.X, pos.Y)].Blocked = false
			pos.X += x
			pos.Y += y

			level.Tiles[index].Blocked = true
			level.mu.Lock()
			level.PlayerVisible.Compute(level, pos.X, pos.Y, 8)
			level.mu.Unlock()
		} else if x != 0 || y != 0 {
			if level.Tiles[index].TileType != WALL {
				monsterPosition := Position{X: pos.X + x, Y: pos.Y + y}
				AttackSystem(g, pos, &monsterPosition)
			}
		}

		updateMapVisibility(level)
	}

	if x != 0 || y != 0 || turnTaken {
		g.Extras.currentX, g.Extras.currentY = 0, 0
		g.Turn = GetNextState(g.Turn)
		g.TurnCounter = 0
	}
}
