package entities

import (
	"github.com/alistairfink/Hussein-Attack/resources"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math"
)

type toiletPaper struct {
	angle  float64
	image  pixel.Picture
	sprite *pixel.Sprite
	win    *pixelgl.Window
	speed  float64
	posX   float64
	posY   float64
}

func NewToiletPaper(resourceLoader *resources.ResourceLoader, win *pixelgl.Window, angle float64) toiletPaper {
	obj := toiletPaper{}
	obj.angle = angle
	obj.image = (*resourceLoader).LoadToiletPaper()
	obj.sprite = pixel.NewSprite(obj.image, obj.image.Bounds())
	obj.win = win
	obj.speed = 5.0

	return obj
}

func (this *toiletPaper) Draw() {
	this.posX += this.speed * math.Cos(this.angle)
	this.posY += this.speed * math.Sin(this.angle)

	// bounds := this.win.Bounds()
	// if this.posX > bounds.Max.X/2+50 ||
	// 	this.posX < -bounds.Max.X/2-50 ||
	// 	this.posY > bounds.Max.Y/2+50 ||
	// 	this.posY < -bounds.Max.Y/2-50 {
	// 	return true
	// }

	matrix := pixel.IM
	matrix = matrix.Moved(pixel.V(100, 100))
	matrix = matrix.ScaledXY(this.win.Bounds().Center(), pixel.V(0.4, 0.4))
	// matrix = matrix.Rotated(this.win.Bounds().Center(), this.angle)
	// matrix = matrix.Moved(pixel.V(this.posX, this.posY))

	this.sprite.Draw(this.win, matrix)

	// return false
}
