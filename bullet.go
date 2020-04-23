package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 2
	bulletSpeed = 0.9
)

type bullet struct {
	tex    *sdl.Texture
	x, y   float64
	angle  float64
	active bool
}

func newBullet(renderer *sdl.Renderer) bullet {
	var b bullet
	b.tex = textureFromBMP(renderer, "sprites/bullet.bmp")
	b.active = false
	return b
}

func (b *bullet) update() {
	b.x += bulletSpeed * math.Cos(b.angle)
	b.y += bulletSpeed * math.Sin(b.angle)

	if b.x > screenWidth || b.x < 0 || b.y > screenHeight || b.y < 0 {
		b.active = false
	}
}

func (b *bullet) draw(renderer *sdl.Renderer) {
	x := b.x - bulletSize/2.0
	y := b.y - bulletSize/2.0

	renderer.Copy(
		b.tex,
		&sdl.Rect{X: 0, Y: 0, W: 1, H: 2},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 2, H: 4},
	)
}

var bulletPool []*bullet

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		b := newBullet(renderer)
		bulletPool = append(bulletPool, &b)
	}
}

func bulletFromPool() (*bullet, bool) {
	for _, bullet := range bulletPool {
		if !bullet.active {
			return bullet, true
		}
	}

	return nil, false
}
