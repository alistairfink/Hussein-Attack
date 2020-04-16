package entities

import (
	"github.com/alistairfink/Hussein-Attack/resources"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
)

type ToiletPaper struct {
	image          pixel.Picture
	sprite         *pixel.Sprite
	win            *pixelgl.Window
	speed          float64
	posX           float64
	posY           float64
	unitStepX      float64
	unitStepY      float64
	angle          float64
	angleIncrement float64
}

func NewToiletPaper(resourceLoader *resources.ResourceLoader, win *pixelgl.Window) ToiletPaper {
	obj := ToiletPaper{}
	obj.image = (*resourceLoader).LoadToiletPaper()
	obj.sprite = pixel.NewSprite(obj.image, obj.image.Bounds())
	obj.win = win

	// Positional Movement
	max, min := 6.0, 1.0
	obj.speed = min + rand.Float64()*(max-min)
	obj.posX, obj.posY = generateEdgeSpawn(win)
	obj.unitStepX, obj.unitStepY = calculateUnitSteps(pixel.V(obj.posX, obj.posY), win.Bounds().Center())

	//  Rotation
	max, min = 0.1, 0.0
	obj.angle = 0.0
	obj.angleIncrement = min + rand.Float64()*(max-min)

	return obj
}

func (this *ToiletPaper) Draw() bool {
	this.angle += this.angleIncrement
	this.posX += this.unitStepX * this.speed
	this.posY += this.unitStepY * this.speed

	centerX, centerY := this.win.Bounds().Center().X, this.win.Bounds().Center().Y
	if abs(centerX-this.posX) < 10 && abs(centerY-this.posY) < 10 {
		return true
	}

	matrix := pixel.IM
	matrix = matrix.Rotated(pixel.V(0.0, 0.0), this.angle)
	matrix = matrix.ScaledXY(pixel.V(0.0, 0.0), pixel.V(0.5, 0.5))
	matrix = matrix.Moved(pixel.V(this.posX, this.posY))

	this.sprite.Draw(this.win, matrix)
	return false
}
