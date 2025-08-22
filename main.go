package main

import (
	"log/slog"
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
	a := app.App(1280, 800, "Roublard")

	scene := core.NewNode()

	cam := camera.New(1)
	cam.SetPosition(0, 0, 3)
	camera.NewOrbitControl(cam)

	g := NewGame(a, scene, cam, slog.Default())

	a.Run(g.Update)
}

type Game struct {
	app   *app.Application
	scene *core.Node
	cam   *camera.Camera

	log *slog.Logger
}

func NewGame(app *app.Application, scene *core.Node, cam *camera.Camera, log *slog.Logger) *Game {
	const tileSideLength = 1.5

	// create a tile
	geom := geometry.NewBox(tileSideLength, .1, tileSideLength)
	mat := material.NewStandard(math32.NewColor("DarkBlue"))
	mesh := graphic.NewMesh(geom, mat)

	// create a button
	btn := gui.NewButton("Make Red")
	btn.SetPosition(100, 40)
	btn.SetSize(40, 40)
	btn.Subscribe(gui.OnClick, func(name string, ev any) {
		mat.SetColor(math32.NewColor("DarkRed"))
	})

	app.Gls().ClearColor(.5, .5, .5, 1)

	onResize := func(evname string, ev any) {
		w, h := app.GetSize()
		app.Gls().Viewport(0, 0, int32(w), int32(h))
		cam.SetAspect(float32(w) / float32(h))
	}

	app.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	// add a point light
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 5)
	pointLight.SetPosition(1, 0, 2)

	gui.Manager().Set(scene)

	// add everything to the scene
	scene.Add(btn)
	scene.Add(light.NewAmbient(&math32.Color{1, 1, 1}, .8))
	scene.Add(cam)
	scene.Add(mesh)
	scene.Add(pointLight)
	scene.Add(helper.NewAxes(0.5))

	return &Game{
		app:   app,
		scene: scene,
		cam:   cam,
		log:   log,
	}
}

func (g *Game) Update(renderer *renderer.Renderer, deltaTime time.Duration) {
	log := g.log.With("func", "update")
	g.app.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)

	err := renderer.Render(g.scene, g.cam)
	if err != nil {
		log.Error("render", "error", err)
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return 1280, 800
}
