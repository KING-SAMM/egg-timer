package main

import (
	"os"
	"gioui.org/app"
	"gioui.org/unit"
	"log"
)


func main() {
	go func() {
	  // create new window
	  w := app.NewWindow(
		app.Title("Egg timer"),
		app.Size(unit.Dp(400), unit.Dp(600)),
	  )
	  if err := draw(w); err != nil {
		log.Fatal(err)
	  }
	  os.Exit(0)
	}()
	app.Main()
  }