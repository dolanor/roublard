package main

import (
	"time"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
)

func main() {
	a := app.App()
	scene := core.NewNode()

	gui.Manager().Set(scene)

	cam := camera.New(1)
	cam.SetPosition(0, 0, 3)

	scene.Add(cam)

	camera.NewOrbitControl(cam)

	onResize := func(evname string, ev any) {
		w, h := a.GetSize()
		a.Gls().Viewport(0, 0, int32(w), int32(h))
		cam.SetAspect(float32(w) / float32(h))
	}

	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	geom := geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
	mat := material.NewStandard(math32.NewColor("DarkBlue"))
	mesh := graphic.NewMesh(geom, mat)
	scene.Add(mesh)

	btn := gui.NewButton("Make Red")
	btn.SetPosition(100, 40)
	btn.SetSize(40, 40)
	btn.Subscribe(gui.OnClick, func(name string, ev any) {
		mat.SetColor(math32.NewColor("DarkRed"))
	})
	scene.Add(btn)

	scene.Add(light.NewAmbient(&math32.Color{1, 1, 1}, .8))
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 5)
	pointLight.SetPosition(1, 0, 2)
	scene.Add(pointLight)

	scene.Add(helper.NewAxes(0.5))

	a.Gls().ClearColor(.5, .5, .5, 1)

	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)
	})
}

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(w, h int) (int, int) {
	return 1280, 800
}
