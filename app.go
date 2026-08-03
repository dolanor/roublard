package main

import (
	"log/slog"
	"time"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/window"
)

type G3NApp struct {
	game *Game

	app         *app.Application
	scene       *core.Node
	cam         *camera.Camera
	orthoToggle bool

	log *slog.Logger
}

func NewG3NApp(log *slog.Logger, game *Game, meshes []core.INode) *G3NApp {
	a := app.App(1280, 800, "Roublard")
	a.SetFullScreen(true)

	scene := core.NewNode()

	cam := camera.New(0)
	cam.SetPosition(40, 50, 25)
	cam.LookAt(&math32.Vector3{40, 0, 25}, &math32.Vector3{0, 0, -1})

	ctl := camera.NewOrbitControl(cam)
	ctl.SetTarget(math32.Vector3{40, 0, 25})

	//a.Gls().ClearColor(.5, .5, .5, 1)
	a.Gls().ClearColor(0, 0, 0, 1)

	onResize := func(_ string, _ any) {
		w, h := a.GetSize()
		a.Gls().Viewport(0, 0, int32(w), int32(h))
		cam.SetAspect(float32(w) / float32(h))
	}

	a.Subscribe(window.OnWindowSize, onResize)
	// force the resize
	onResize("", nil)

	gui.Manager().Set(scene)

	scene.Add(cam)

	// Add the level meshes
	for _, t := range game.Map.CurrentLevel.Tiles {
		scene.Add(t.Mesh)
	}

	// Add the characters meshes
	for _, mesh := range meshes {
		scene.Add(mesh)
	}

	return &G3NApp{
		game: game,

		app:   a,
		scene: scene,
		cam:   cam,
		log:   log,
	}
}

func (g3nApp *G3NApp) UpdateLogicLoop() {
	for range time.Tick(time.Second / 120) {
		err := g3nApp.UpdateLogic()
		if err != nil {
			g3nApp.log.Error("update logic", "error", err)
		}
	}
}

// Update is called each tic.
func (g3nApp *G3NApp) UpdateLogic() error {
	g3nApp.game.TurnCounter++
	if g3nApp.game.Turn == PlayerTurn && g3nApp.game.TurnCounter > 20 {
		TakePlayerAction(g3nApp.game)
	}
	if g3nApp.game.Turn == MonsterTurn {
		UpdateMonster(g3nApp.game)
	}

	ProcessUserLogG3N(g3nApp.game)
	ProcessHUDG3N(g3nApp.game)

	return nil

}

// Update is the rendering update callback for g3n.
// It is different from the Update() callback for ebiten which is more the logic update callback.
// our logic callback is [G3NApp.UpdateLogicLoop].
func (g3nApp *G3NApp) Update(renderer *renderer.Renderer, deltaTime time.Duration) {
	log := g3nApp.log.With("func", "update")
	g3nApp.app.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)

	ProcessRenderables(g3nApp.game, g3nApp.game.Map.CurrentLevel)

	err := renderer.Render(g3nApp.scene, g3nApp.cam)
	if err != nil {
		log.Error("render", "error", err)
	}
}

func (g3nApp *G3NApp) SetupKeyboardEventHandlers() {
	g3nApp.app.Subscribe(window.OnKeyDown, g3nApp.onKey)
}

func (g3nApp *G3NApp) GetSize() (width, height int) {
	return g3nApp.app.GetSize()
}

func (g3nApp *G3NApp) SetupUI() {
	// we need to sleep a little bit otherwise the app.GetSize() will be wrong
	time.Sleep(100 * time.Millisecond)

	width, height := g3nApp.GetSize()
	slog.Info("size", "w", width, "h", height)
	panel := CreatePanel(g3nApp.scene, width, height)

	CreateLogWindow(panel, width/2, height/4)
	CreateHUDWindow(panel, width/2, height/4)
}

func (g3nApp *G3NApp) Run() {
	go g3nApp.UpdateLogicLoop()
	g3nApp.app.Run(g3nApp.Update)
}
