package main

import (
	"math"
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

	return &player
}

func (p *player) draw(renderer *sdl.Renderer) {
	// converting player coords to middle of sprite
	x := p.x - playerSize/2.0
	y := p.y - playerSize/2.0

	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 16, H: 16},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 32, H: 32},
	)
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		// move player left
		if p.x-(playerSize/2.0) > 0 {
			p.x -= playerSpeed
		}

	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		// move player right
		if p.x+(playerSize/2.0)+15 < screenWidth {
			p.x += playerSpeed
		}
	}

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastShot) >= playerShotCooldown {
			// TODO: maybe update these magic offset numbers
			p.shoot(p.x+7, p.y-10)
		}
	}
}

func (p *player) shoot(x, y float64) {
	if b, ok := bulletFromPool(); ok {
		b.active = true
		b.x = x
		b.y = y
		b.angle = 270 * (math.Pi / 180)

		p.lastShot = time.Now()
	}
}
