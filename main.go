package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing sdl: ", err)
		return
	}

	window, err := sdl.CreateWindow("gaming in go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		600, 800, sdl.WINDOW_SHOWN,
	)

	if err != nil {
		fmt.Println("initializing window: ", err)
		return
	}

	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if err != nil {
		fmt.Println("creating renderer: ", err)
		return
	}

	defer renderer.Destroy()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()
		renderer.Present()
	}
}
