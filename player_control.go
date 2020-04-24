package main

import "github.com/veandco/go-sdl2/sdl"

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
