package main

import (
	"gioui.org/app"
)

func main() {
	go func(){
		// Create a new window 
		w := app.NewWindow()

		// Listen for events on the window 
		for range w.Events() {
			
		}
	}()

	app.Main()
}