package main

import (
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/window"

	"github.com/dolanor/roublard/assets"
)

type Move string

const (
	MoveLeft  Move = "move_left"
	MoveRight Move = "move_right"
	MoveUp    Move = "move_up"
	MoveDown  Move = "move_down"
)

func (g3nApp *G3NApp) updateMapVisibility(level *Level) {
	gd := NewGameData()
	solidMat := g3nApp.materialManager.Get(assets.MaterialID("wall"))
	wireframeMat := g3nApp.materialManager.Get(assets.MaterialID("wallwf"))
	_, _ = solidMat, wireframeMat
	// We decide to check for every tile in the level if it should be rendered or not
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < levelHeight; y++ {
			index := level.GetIndexFromXY(x, y)
			tile := level.Tiles[index]

			level.mu.Lock()
			isVisible := level.PlayerVisible.IsVisible(x, y)
			level.mu.Unlock()

			if isVisible {
				tile.IsRevealed = true

				tile.Mesh.SetVisible(true)
				if tile.IsWall {
					tile.Mesh.SetMaterial(solidMat)
				}
			} else {
				if !tile.IsRevealed {
					tile.Mesh.SetVisible(false)
					continue
				}

				if tile.IsWall {
					tile.Mesh.SetMaterial(wireframeMat)
					tile.Mesh.SetVisible(true)
					continue
				}
			}
		}
	}
}

func (g3nApp *G3NApp) onKey(evname string, ev any) {
	g3nApp.game.currentX, g3nApp.game.currentY = g3nApp.processKeys(ev)
}

func (g3nApp *G3NApp) processKeys(ev any) (x, y int) {

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
			g3nApp.app.Exit()
			return x, y
		}
		// should deal with turn taken to be full iso compliant, but I don't think it matters that much and I
		// don't want to deal with this event logic now

	case window.KeyU:
		if kev.Mods == window.ModControl {
			if g3nApp.orthoToggle {
				g3nApp.cam.SetProjection(camera.Orthographic)
				g3nApp.orthoToggle = !g3nApp.orthoToggle
			} else {
				g3nApp.cam.SetProjection(camera.Perspective)
				g3nApp.orthoToggle = !g3nApp.orthoToggle
			}
		}
	case window.KeySlash:
		fs := g3nApp.app.FullScreen()
		g3nApp.app.SetFullScreen(!fs)
		return x, y
	}

	return x, y
}
