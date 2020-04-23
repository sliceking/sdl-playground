package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	enemySpeed = 0.3
	enemySize  = 16
)

type enemy struct {
	tex  *sdl.Texture
	x, y float64
}

func newEnemy(renderer *sdl.Renderer, x, y float64) (enemy, error) {
	var e enemy
	e.tex = textureFromBMP(renderer, "sprites/baddie.bmp")
	e.x = x
	e.y = y

	return e, nil
}

func (e *enemy) draw(renderer *sdl.Renderer) {
	// converting enemy coords to middle of sprite
	x := e.x - enemySize/2.0
	y := e.y - enemySize/2.0

	renderer.CopyEx(
		e.tex,
		&sdl.Rect{X: 0, Y: 0, W: 16, H: 16},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 32, H: 32},
		180,
		&sdl.Point{X: enemySize / 2, Y: enemySize / 2},
		sdl.FLIP_NONE,
	)
}

func (e *enemy) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		// move enemy left
		e.x -= enemySpeed

	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		// move enemy right
		e.x += enemySpeed
	}
}
