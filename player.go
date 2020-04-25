package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 0.3
	playerSize         = 16
	playerShotCooldown = time.Millisecond * 250
)

type player struct {
	tex      *sdl.Texture
	x, y     float64
	lastShot time.Time
}

func newPlayer(renderer *sdl.Renderer) *element {
	player := element{}
	player.position = vector{
		x: screenWidth / 2.0,
		y: screenHeight - playerSize*2.5,
	}
	player.active = true

	sr := newSpriteRenderer(&player, renderer, "sprites/hero.bmp")
	player.addComponent(sr)

	mover := newKeyboardMover(&player, playerSpeed)
	player.addComponent(mover)

	shooter := newKeyboardShooter(&player, playerShotCooldown)
	player.addComponent(shooter)

	return &player
}
