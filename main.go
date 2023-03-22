package main

import (
	"os"
	"gioui.org/app"
	"gioui.org/unit"
	"log"
	"time"
)

// Define the progress variables, a channel and a variable
var progressIncrementer chan float32
var progress float32

func main() {
	// Setup a separate channel to provide ticks to increment progress
	progressIncrementer = make(chan float32)
	go func() {
	  for {
		time.Sleep(time.Second / 25)
		progressIncrementer <- 0.004
	  }
	}()

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