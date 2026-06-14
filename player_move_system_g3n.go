package main

import (
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/window"

	"github.com/dolanor/roublard/assets"
)

var ebiten ebitenFake

type ebitenFake struct {
	KeyUp    string
	KeyDown  string
	KeyRight string
	KeyLeft  string
	KeyQ     string
}

func (ebitenFake) IsKeyPressed(key string) bool {
	return false
}

func updateMapVisibility(level *Level) {
	gd := NewGameData()
	solidMat := level.mm.Get(assets.MaterialID("wall"))
	wireframeMat := level.mm.Get(assets.MaterialID("wallwf"))
	_, _ = solidMat, wireframeMat
	// We decide to check for every tile in the level if it should be rendered or not
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			index := level.GetIndexFromXY(x, y)
			tile := level.Tiles[index]

			level.mu.Lock()
			isVisible := level.PlayerVisible.IsVisible(x, y)
			level.mu.Unlock()

			if isVisible {
				tile.IsRevealed = true

				tile.Image.SetVisible(true)
				if tile.IsWall {
					tile.Image.SetMaterial(solidMat)
				}
			} else {
				if !tile.IsRevealed {
					tile.Image.SetVisible(false)
					continue
				}

				if tile.IsWall {
					tile.Image.SetMaterial(wireframeMat)
					tile.Image.SetVisible(true)
					continue
				}
			}
		}
	}
}

type Move string

const (
	MoveLeft  Move = "move_left"
	MoveRight Move = "move_right"
	MoveUp    Move = "move_up"
	MoveDown  Move = "move_down"
)

func (g *Game) onKey(evname string, ev any) {
	g.Extras.currentX, g.Extras.currentY = g.processKeys(ev)
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
			g.Extras.app.Exit()
			return x, y
		}
		// should deal with turn taken to be full iso compliant, but I don't think it matters that much and I
		// don't want to deal with this event logic now

	case window.KeyU:
		if kev.Mods == window.ModControl {
			if g.Extras.orthoToggle {
				g.Extras.cam.SetProjection(camera.Orthographic)
				g.Extras.orthoToggle = !g.Extras.orthoToggle
			} else {
				g.Extras.cam.SetProjection(camera.Perspective)
				g.Extras.orthoToggle = !g.Extras.orthoToggle
			}

		}
	}

	return x, y
}
