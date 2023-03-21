package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)


func draw(w *app.Window) error {
	// ops are the operations for the unit
	var ops op.Ops
	
	// startButton is a clickable widget
	var startButton widget.Clickable
	
	// th defines the material design style
	th := material.NewTheme(gofont.Collection())

  // listen for events in the window.
  for e := range w.Events() {

    // detect what type of event
    switch e := e.(type) {

    // this is sent when the application should re-render.
    case system.FrameEvent:
        gtx := layout.NewContext(&ops, e)

			// Flexbox layout
			layout.Flex{
				// Vertical alignment, from top to bottom
				Axis: layout.Vertical,

				// Empty space is lefr at the start, i.e at the top
				Spacing: layout.SpaceStart,
			}.Layout(
				gtx,

				// We insert two rigid elements:
				// First a button ...
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &startButton, "Start")
						return btn.Layout(gtx)
					},
				),

				// ..then an empty spacer
				layout.Rigid(
					// The height of the spacer is 25 Device independent pixels
					layout.Spacer{ Height: unit.Dp(25)}.Layout,
				),
			)
			e.Frame(gtx.Ops)

    // this is sent when the application is closed.
    case system.DestroyEvent:
      return e.Err
    }
  }
  return nil
}