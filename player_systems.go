package main

func (g3nApp *G3NApp) TakePlayerAction() {
	// TODO: rename g -> game
	g := g3nApp.game

	players := g.WorldTags["players"]
	turnTaken := false

	x := g.currentX
	y := g.currentY

	level := g.Map.CurrentLevel

	for _, result := range g.World.Query(players) {
		pos := result.Components[positions].(*Position)
		index := level.GetIndexFromXY(pos.X+x, pos.Y+y)

		tile := level.Tiles[index]
		if !tile.Blocked {
			level.Tiles[level.GetIndexFromXY(pos.X, pos.Y)].Blocked = false
			pos.X += x
			pos.Y += y

			level.Tiles[index].Blocked = true
			level.mu.Lock()
			level.PlayerVisible.Compute(level, pos.X, pos.Y, 8)
			level.mu.Unlock()
		} else if x != 0 || y != 0 {
			if level.Tiles[index].TileType != TileTypeWall {
				//Its a tile with a monster -- Fight it
				monsterPosition := Position{X: pos.X + x, Y: pos.Y + y}

				AttackSystem(g, pos, &monsterPosition)
			}
		}

		g3nApp.updateMapVisibility(level)
	}

	if x != 0 || y != 0 || turnTaken {
		g.currentX, g.currentY = 0, 0
		g.Turn = GetNextState(g.Turn)
		g.TurnCounter = 0
	}
}
