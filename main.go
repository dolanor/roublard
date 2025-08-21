package main

import (
	"log/slog"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Roublard")

	err := ebiten.RunGame(g)
	if err != nil {
		slog.Error("run game", "error", err)
		os.Exit(1)
	}
}

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(w, h int) (int, int) {
	return 1280, 800
}
