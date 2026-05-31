package main

import (
	"log/slog"
	"time"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/window"
)

func (g *Game) UpdateLogicLoop() {
	for range time.Tick(time.Second / 120) {
		err := g.UpdateLogic()
		if err != nil {
			g.Extras.log.Error("update logic", "error", err)
		}
	}
}

// Update is the rendering update callback for g3n.
// It is different from the Update() callback for ebiten which is more the logic update callback.
// our logic callback is [Game.logicUpdateLoop].
func (g *Game) Update(renderer *renderer.Renderer, deltaTime time.Duration) {
	log := g.Extras.log.With("func", "update")
	g.Extras.app.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)

	ProcessRenderables(g, g.Map.CurrentLevel)

	err := renderer.Render(g.Extras.scene, g.Extras.cam)
	if err != nil {
		log.Error("render", "error", err)
	}
}

type GameExtras struct {
	// FIXME: maybe protect it with a mutex
	currentX int
	currentY int

	app         *app.Application
	scene       *core.Node
	cam         *camera.Camera
	orthoToggle bool

	log *slog.Logger
}

func NewG3NExtras() *GameExtras {
	a := app.App(1280, 800, "Roublard")
	a.IWindow.SetFullScreen(true)

	scene := core.NewNode()

	cam := camera.New(0)
	cam.SetPosition(40, 50, 25)
	cam.LookAt(&math32.Vector3{40, 0, 25}, &math32.Vector3{0, 0, -1})

	ctl := camera.NewOrbitControl(cam)
	ctl.SetTarget(math32.Vector3{40, 0, 25})

	a.Gls().ClearColor(.5, .5, .5, 1)
	//app.Gls().ClearColor(0, 0, 0, 1)

	onResize := func(_ string, _ any) {
		w, h := a.GetSize()
		a.Gls().Viewport(0, 0, int32(w), int32(h))
		cam.SetAspect(float32(w) / float32(h))
	}

	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	// add a point light
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 30)
	pointLight.SetPosition(1, 1, 2)

	gui.Manager().Set(scene)

	// add everything to the scene
	//scene.Add(light.NewAmbient(&math32.Color{1, 1, 1}, .8))
	scene.Add(pointLight)
	scene.Add(cam)

	return &GameExtras{
		app:   a,
		scene: scene,
		cam:   cam,
		log:   slog.Default(),
	}
}
