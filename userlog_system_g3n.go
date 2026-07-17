package main

import (
	"fmt"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
)

var logList *gui.List

func (g *Game) CreateLogWindow() {
	dl := gui.NewDockLayout()

	width, height := g.Extras.app.GetSize()

	panel := gui.NewPanel(float32(width), float32(height))
	panel.SetColor4(&gui.StyleDefault().Scroller.BgColor)
	panel.SetLayout(dl)
	panel.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockCenter})
	panel.SetRenderable(false)
	panel.SetEnabled(false)

	w1 := gui.NewWindow(float32(width)/2, float32(height)/4)
	w1.SetPosition(10, float32(height)*3/4)
	w1.SetResizable(true)

	logList = gui.NewVList(w1.Width(), w1.Height())

	w1.Add(logList)
	panel.Add(w1)

	g.Extras.scene.Add(panel)
	gui.Manager().Set(panel)
}

func ProcessUserLogG3N(g *Game) {
	currentMessages := make([]string, 0, 5)

	for _, m := range g.World.Query(g.WorldTags["messengers"]) {
		messages := m.Components[userMessage].(*UserMessage)
		if messages.AttackMessage != "" {
			currentMessages = append(currentMessages, messages.AttackMessage)
			fmt.Print(messages.AttackMessage)
			messages.AttackMessage = ""
		}
	}
	for _, m := range g.World.Query(g.WorldTags["messengers"]) {
		messages := m.Components[userMessage].(*UserMessage)
		if messages.DeadMessage != "" {
			currentMessages = append(currentMessages, messages.DeadMessage)
			fmt.Print(messages.DeadMessage)
			messages.DeadMessage = ""
			g.World.DisposeEntity(m.Entity)
		}
		if messages.GameStateMessage != "" {
			currentMessages = append(currentMessages, messages.GameStateMessage)
			fmt.Print(messages.GameStateMessage)
			messages.GameStateMessage = ""
		}

	}

	for _, msg := range currentMessages {
		if msg != "" {
			// TODO: remove that later
			// remove the \n that is currently added in the combat system
			msg := msg[:len(msg)-1]
			lbl := gui.NewLabel(msg)
			lbl.SetColor(math32.NewColor("black"))
			lbl.SetBgColor(math32.NewColor("white"))
			logList.Add(lbl)
			logList.ScrollDown()
			logList.ScrollDown()
		}
	}
}
