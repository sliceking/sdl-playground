package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 2
	bulletSpeed = 0.9
)

func newBullet(renderer *sdl.Renderer) *element {
	bullet := element{}

	sr := newSpriteRenderer(&bullet, renderer, "sprites/bullet.bmp")
	bullet.addComponent(sr)

	mover := newBulletMover(&bullet, bulletSpeed)
	bullet.addComponent(mover)

	bullet.active = false

	return &bullet
}

var bulletPool []*element

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		b := newBullet(renderer)
		elements = append(elements, b)
		bulletPool = append(bulletPool, b)
	}
}

func bulletFromPool() (*element, bool) {
	for _, bullet := range bulletPool {
		if !bullet.active {
			return bullet, true
		}
	}

	return nil, false
}
