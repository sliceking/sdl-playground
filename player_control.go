package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element
	speed     float64

	sr *spriteRenderer
}

func newKeyboardMover(container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()

	cont := mover.container

	if keys[sdl.SCANCODE_LEFT] == 1 {
		// move left
		if cont.position.x-(mover.sr.width/2.0) > 0 {
			cont.position.x -= mover.speed
		}

	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		// move right
		if cont.position.x+(mover.sr.height/2.0)+15 < screenWidth {
			cont.position.x += mover.speed
		}
	}

	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardMover) onCollision(other *element) error {
	return nil
}

type keyboardShooter struct {
	container *element
	cooldown  time.Duration
	lastShot  time.Time
}

func newKeyboardShooter(container *element, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		cooldown:  cooldown,
	}
}

func (s *keyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()

	pos := s.container.position

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(s.lastShot) >= s.cooldown {
			// TODO: maybe update these magic offset numbers
			s.shoot(pos.x+7, pos.y-10)
		}
	}

	return nil
}

func (s *keyboardShooter) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (s *keyboardShooter) onCollision(other *element) error {
	return nil
}

func (s *keyboardShooter) shoot(x, y float64) {
	if b, ok := bulletFromPool(); ok {
		b.active = true
		b.position.x = x
		b.position.y = y
		b.rotation = 270 * (math.Pi / 180)

		s.lastShot = time.Now()
	}
}
