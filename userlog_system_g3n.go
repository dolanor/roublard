package main

import (
	"fmt"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
)

var logList *gui.List

func CreatePanel(scene *core.Node, width, height int) *gui.Panel {
	dl := gui.NewDockLayout()
	panel := gui.NewPanel(float32(width), float32(height))
	panel.SetColor4(&gui.StyleDefault().Scroller.BgColor)
	panel.SetLayout(dl)
	panel.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockCenter})
	panel.SetRenderable(false)
	panel.SetEnabled(false)

	scene.Add(panel)
	gui.Manager().Set(panel)

	return panel
}

func CreateLogWindow(panel *gui.Panel, width, height int) {
	w1 := gui.NewWindow(float32(width), float32(height))
	w1.SetPosition(0, 3*float32(height))
	w1.SetResizable(true)

	logList = gui.NewVList(w1.Width(), w1.Height())

	w1.Add(logList)
	panel.Add(w1)
}

func ProcessUserLogG3N(g *Game) {
	currentMessages := make([]string, 0, 5)

	for _, m := range g.World.Query(g.WorldTags["messengers"]) {
		messages := m.Components[userMessages].(*UserMessage)
		if messages.AttackMessage != "" {
			currentMessages = append(currentMessages, messages.AttackMessage)
			fmt.Println(messages.AttackMessage)
			messages.AttackMessage = ""
		}
	}
	for _, m := range g.World.Query(g.WorldTags["messengers"]) {
		messages := m.Components[userMessages].(*UserMessage)
		if messages.DeadMessage != "" {
			currentMessages = append(currentMessages, messages.DeadMessage)
			fmt.Println(messages.DeadMessage)
			messages.DeadMessage = ""
			g.World.DisposeEntity(m.Entity)
		}
		if messages.GameStateMessage != "" {
			currentMessages = append(currentMessages, messages.GameStateMessage)
			fmt.Println(messages.GameStateMessage)
			messages.GameStateMessage = ""
		}

	}

	for _, msg := range currentMessages {
		if msg != "" {
			lbl := gui.NewLabel(msg)
			lbl.SetColor(math32.NewColor("black"))
			lbl.SetBgColor(math32.NewColor("white"))
			logList.Add(lbl)
			logList.ScrollDown()
			logList.ScrollDown()
		}
	}
}
