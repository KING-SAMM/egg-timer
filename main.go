package main

import (
	"os"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func(){
		// Create a new window 
		w := app.NewWindow(
			app.Title("Egg Timer"),
			app.Size(unit.Dp(400), unit.Dp(600)),
		)

		// ops are the operations for the unit
		var ops op.Ops

		// startButton is a clickable widget
		var startButton widget.Clickable

		// th defines the material design style
		th := material.NewTheme(gofont.Collection())

		// Listen for events in the window 
		for e :=  range w.Events() {

			// detect what type of event
			switch e := e.(type) {

				//This is sent when the application should re-render
				case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				btn := material.Button(th, &startButton, "Start")
				btn.Layout(gtx)
				e.Frame(gtx.Ops)
			}
		}
		// Ensures cleaner exit
		os.Exit(0)
	}()
	app.Main()
}