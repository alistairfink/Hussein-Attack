package entities

import (
	"github.com/alistairfink/Hussein-Attack/resources"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
)

type toiletPaper struct {
	angle     float64
	image     pixel.Picture
	sprite    *pixel.Sprite
	win       *pixelgl.Window
	speed     float64
	posX      float64
	posY      float64
	unitStepX float64
	unitStepY float64
}

func NewToiletPaper(resourceLoader *resources.ResourceLoader, win *pixelgl.Window) toiletPaper {
	obj := toiletPaper{}
	obj.image = (*resourceLoader).LoadToiletPaper()
	obj.sprite = pixel.NewSprite(obj.image, obj.image.Bounds())
	obj.win = win
	max, min := 6.0, 1.0
	obj.speed = min + rand.Float64()*(max-min)
	obj.posX, obj.posY = generateEdgeSpawn(win)
	obj.angle = 0.0
	obj.unitStepX, obj.unitStepY = calculateUnitSteps(pixel.V(obj.posX, obj.posY), win.Bounds().Center())

	return obj
}

func (this *toiletPaper) Draw() {
	this.angle += 0.05
	this.posX += this.unitStepX * this.speed
	this.posY += this.unitStepY * this.speed

	matrix := pixel.IM
	// matrix = matrix.ScaledXY(pixel.V(this.posX, this.posY), pixel.V(0.4, 0.4))
	matrix = matrix.Moved(pixel.V(this.posX, this.posY))
	// matrix = matrix.Moved(this.win.Bounds().Center())

	this.sprite.Draw(this.win, matrix)
}
