package entities

import (
	"github.com/alistairfink/Hussein-Attack/resources"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math"
)

type laser struct {
	angle  float64
	image  pixel.Picture
	sprite *pixel.Sprite
	win    *pixelgl.Window
	speed  float64
	posX   float64
	posY   float64
}

func NewLaser(resourceLoader *resources.ResourceLoader, win *pixelgl.Window, angle float64) laser {
	obj := laser{}
	obj.angle = angle
	obj.image = (*resourceLoader).LoadLaser()
	obj.sprite = pixel.NewSprite(obj.image, obj.image.Bounds())
	obj.win = win
	obj.speed = 3

	return obj
}

func (this *laser) Draw() bool {
	this.posX += this.speed * math.Cos(this.angle)
	this.posY += this.speed * math.Sin(this.angle)

	bounds := this.win.Bounds()
	if this.posX > bounds.Max.X/2+50 ||
		this.posX < -bounds.Max.X/2-50 ||
		this.posY > bounds.Max.Y/2+50 ||
		this.posY < -bounds.Max.Y/2-50 {
		return true
	}

	matrix := pixel.IM
	matrix = matrix.Moved(this.win.Bounds().Center())
	matrix = matrix.ScaledXY(this.win.Bounds().Center(), pixel.V(0.5, 0.5))
	matrix = matrix.Rotated(this.win.Bounds().Center(), this.angle)
	matrix = matrix.Moved(pixel.V(this.posX, this.posY))

	this.sprite.Draw(this.win, matrix)

	return false
}
