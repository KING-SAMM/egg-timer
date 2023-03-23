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
	"image"
	"image/color"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

// Constraints and Dimensions
type C = layout.Context
type D = layout.Dimensions

type Point struct {
    X, Y float32
}

type Rectangle struct {
    Min, Max Point
}

func draw(w *app.Window) error {
	// ops are the operations for the unit
	var ops op.Ops
	
	// startButton is a clickable widget
	var startButton widget.Clickable
	
	// th defines the material design style
	th := material.NewTheme(gofont.Collection())

	// is the egg boiling?
	var boiling bool

	for {
		select {
			// Listen for events in the window
			case e := <-w.Events():
				// detect what type of event
				switch e := e.(type) {

					// this is sent when the application should re-render.
					case system.FrameEvent:
						gtx := layout.NewContext(&ops, e)

						// Let's try out the flexbox layout concept
						if startButton.Clicked() {
							boiling = !boiling
						}
			
						// Flexbox layout
						layout.Flex{
							// Vertical alignment, from top to bottom
							Axis: layout.Vertical,
			
							// Empty space is lefr at the start, i.e at the top
							Spacing: layout.SpaceStart,
						}.Layout(
							gtx,

							// Egg as Cirle
							layout.Rigid(
								func(gtx C) D {
									circle := clip.Ellipse{
										// Hard coding the x coordinate. Try resizing the window
										Min: image.Pt(80, 0),
										Max: image.Pt(320, 240),
										// Soft coding the x coordinate. Try resizing the window
										//Min: image.Pt(gtx.Constraints.Max.X/2-120, 0),
										//Max: image.Pt(gtx.Constraints.Max.X/2+120, 240),
									}.Op(gtx.Ops)
									color := color.NRGBA{R: 200, A: 255}
									paint.FillShape(gtx.Ops, color, circle)
									d := image.Point{Y: 400}
									return layout.Dimensions{Size: d}
								},
							),
			
							// Progressbar
							layout.Rigid(
								func(gtx C) D {
								bar := material.ProgressBar(th, progress)  // Here progress is used
								return bar.Layout(gtx)
								},
							),
			
							// Button
							layout.Rigid(
								func(gtx C) D {
									// ONE: First define margins around the button using layout.Inset ...
									margins := layout.Inset{
										Top:    unit.Dp(25),
										Bottom: unit.Dp(25),
										Right:  unit.Dp(35),
										Left:   unit.Dp(35),
									}
									// TWO: ...then we lay out those margins...
									return margins.Layout(
										gtx,
										// THREE: ...and finally within the margins, we define and lay out the button
										func(gtx C) D {
											var text string
											if !boiling {
												text = "Start"
											} else {
												text = "Stop"
											}
			
											btn := material.Button(th, &startButton, text)
											return btn.Layout(gtx)
										},
									)
								},
							),
						)
						e.Frame(gtx.Ops)
			
					// this is sent when the application is closed.
					case system.DestroyEvent:
						return e.Err
				}
				
			// listen for events in the incrementor channel
			case p := <-progressIncrementer:
				if boiling && progress < 1 {
					progress += p
					w.Invalidate()
				}
		}	
	}
}

