package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	tex *sdl.Texture
}

func newPlayer(renderer *sdl.Renderer) (player, error) {
	var p player
	img, err := sdl.LoadBMP("sprites/hero.bmp")
	if err != nil {
		return player{}, fmt.Errorf("loading player bmp: ", err)
	}
	defer img.Free()

	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("creating texture player: ", err)
	}

	return p, nil
}

func (p *player) draw(renderer *sdl.Renderer) {
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 16, H: 16},
		&sdl.Rect{X: 0, Y: 0, W: 32, H: 32},
	)
}
