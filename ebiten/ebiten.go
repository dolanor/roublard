package ebiten

import "github.com/g3n/engine/graphic"

type Image string

func (Image) DrawImage(img *graphic.Mesh, op *DrawImageOptions) {}

type DrawImageOptions struct {
	GeoM   Geom
	ColorM Colorm
}

type Geom struct{}

func (Geom) Translate(x, y float64) {}

type Colorm struct{}

func (Colorm) Translate(x, y, w, h float64) {}

var (
	KeyUp    string
	KeyDown  string
	KeyRight string
	KeyLeft  string
	KeyQ     string
)

func IsKeyPressed(key string) bool {
	return false
}

func SetWindowResizable(bool) {}
func SetWindowTitle(string)   {}
