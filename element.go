package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type vector struct {
	x, y float64
}

type component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
	onCollision(other *element) error
}

type element struct {
	position   vector
	rotation   float64
	active     bool
	collisions []circle
	components []component
}

func (e *element) addComponent(new component) {
	for _, existing := range e.components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf("adding duped component: %+v", reflect.TypeOf(new)))
		}
	}

	e.components = append(e.components, new)
}

func (e *element) getComponent(withType component) component {
	typ := reflect.TypeOf(withType)
	for _, comp := range e.components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}

	panic(fmt.Sprintf("couldnt find component getcomponent: %+v", reflect.TypeOf(withType)))
}

func (e *element) draw(renderer *sdl.Renderer) error {
	for _, comp := range e.components {
		err := comp.onDraw(renderer)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *element) update() error {
	for _, comp := range e.components {
		err := comp.onUpdate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *element) collision(other *element) error {
	for _, comp := range e.components {
		err := comp.onCollision(other)
		if err != nil {
			return err
		}
	}

	return nil
}

var elements []*element
