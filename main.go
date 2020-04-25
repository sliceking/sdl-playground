package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing sdl: ", err)
		return
	}

	window, err := sdl.CreateWindow("gaming in go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight, sdl.WINDOW_OPENGL,
	)

	if err != nil {
		fmt.Println("initializing window: ", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("creating renderer: ", err)
		return
	}
	defer renderer.Destroy()

	player := newPlayer(renderer)

	var enemies []enemy
	// Create our rows of enemies
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i) / 5) * screenWidth
			y := float64(j)*enemySize*2 + (enemySize * 2)

			enemy, err := newEnemy(renderer, x, y)
			if err != nil {
				fmt.Println("enemy creation: ", err)
				return
			}
			enemies = append(enemies, enemy)
		}
	}

	initBulletPool(renderer)

	// Main loop
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		err = player.update()
		if err != nil {
			fmt.Println("updating player", err)
			return
		}

		err = player.draw(renderer)
		if err != nil {
			fmt.Println("drawing player", err)
			return
		}

		for _, enemy := range enemies {
			enemy.draw(renderer)
		}

		for _, bullet := range bulletPool {
			bullet.update()
			bullet.draw(renderer)
		}

		renderer.Present()
	}
}
