package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	enemySpeed = 0.3
	enemySize  = 16
)

func newBasicEnemy(renderer *sdl.Renderer, position vector) *element {
	basicEnemy := element{}

	basicEnemy.position = position
	basicEnemy.rotation = 180

	sr := newSpriteRenderer(&basicEnemy, renderer, "sprites/baddie.bmp")
	basicEnemy.addComponent(sr)

	basicEnemy.active = true

	col := circle{
		center: basicEnemy.position,
		radius: 38,
	}
	basicEnemy.collisions = append(basicEnemy.collisions, col)

	return &basicEnemy
}
