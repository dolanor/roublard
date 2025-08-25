package main

import (
	"log/slog"
	"time"

	"github.com/bytearena/ecs"
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

func main() {
	a := app.App(1280, 800, "Roublard")

	scene := core.NewNode()

	cam := camera.New(0)
	cam.SetPosition(40, 50, 25)
	ctl := camera.NewOrbitControl(cam)
	ctl.SetTarget(math32.Vector3{40, 0, 25})
	//cam.SetDirection(110, 10, -100)

	g := NewGame(a, scene, cam, slog.Default())

	a.Run(g.Update)
}

type Game struct {
	app   *app.Application
	scene *core.Node
	cam   *camera.Camera

	gameMap   GameMap
	World     *ecs.Manager
	WorldTags map[string]ecs.Tag

	log *slog.Logger
}

func NewGame(app *app.Application, scene *core.Node, cam *camera.Camera, log *slog.Logger) *Game {
	app.Gls().ClearColor(.5, .5, .5, 1)

	onResize := func(evname string, ev any) {
		w, h := app.GetSize()
		app.Gls().Viewport(0, 0, int32(w), int32(h))
		cam.SetAspect(float32(w) / float32(h))
	}

	app.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	// add a point light
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 30)
	pointLight.SetPosition(1, 1, 2)

	gui.Manager().Set(scene)
	gm := NewGameMap()
	tiles := gm.Dungeons[0].Levels[0].Tiles
	for _, t := range tiles {
		scene.Add(t.Mesh)
	}

	// add everything to the scene
	//scene.Add(light.NewAmbient(&math32.Color{1, 1, 1}, .8))
	scene.Add(pointLight)
	scene.Add(cam)

	world, tags := InitWorld(scene)

	return &Game{
		app:   app,
		scene: scene,
		cam:   cam,

		gameMap:   gm,
		World:     world,
		WorldTags: tags,

		log: log,
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
	gd := NewGameData()
	return gd.TileWidth * gd.ScreenWidth, gd.TileHeight * gd.ScreenHeight
}
