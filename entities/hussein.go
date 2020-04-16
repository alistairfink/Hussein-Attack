package entities

import (
	"github.com/alistairfink/2D-Game-Fun/resources"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Hussein struct {
	angle       float64
	image       pixel.Picture
	sprite      *pixel.Sprite
	win         *pixelgl.Window
	rotateSpeed float64
}

func NewHussein(resourceLoader *resources.ResourceLoader, win *pixelgl.Window) Hussein {
	obj := Hussein{}
	obj.angle = 0.0
	obj.image = (*resourceLoader).LoadHussein()
	obj.sprite = pixel.NewSprite(obj.image, obj.image.Bounds())
	obj.win = win
	obj.rotateSpeed = 4
	return obj
}

func (this *Hussein) Draw() {
	matrix := pixel.IM
	matrix = matrix.Moved(this.win.Bounds().Center())
	matrix = matrix.ScaledXY(this.win.Bounds().Center(), pixel.V(0.15, 0.15))
	matrix = matrix.Rotated(this.win.Bounds().Center(), this.angle)
	this.sprite.Draw(this.win, matrix)
}

func (this *Hussein) RotateLeft(deltaTime float64) {
	this.angle += deltaTime * this.rotateSpeed
	this.Draw()
}

func (this *Hussein) RotateRight(deltaTime float64) {
	this.angle -= deltaTime * this.rotateSpeed
	this.Draw()
}