package entities

import (
	"github.com/alistairfink/Hussein-Attack/resources"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type hussein struct {
	angle          float64
	image          pixel.Picture
	sprite         *pixel.Sprite
	win            *pixelgl.Window
	rotateSpeed    float64
	lasers         []laser
	resourceLoader *resources.ResourceLoader
}

func NewHussein(resourceLoader *resources.ResourceLoader, win *pixelgl.Window) hussein {
	obj := hussein{}
	obj.angle = 0.0
	obj.image = (*resourceLoader).LoadHussein()
	obj.sprite = pixel.NewSprite(obj.image, obj.image.Bounds())
	obj.win = win
	obj.rotateSpeed = 4
	obj.lasers = []laser{}
	obj.resourceLoader = resourceLoader

	return obj
}

func (this *hussein) Draw() {
	matrix := pixel.IM
	matrix = matrix.Moved(this.win.Bounds().Center())
	matrix = matrix.ScaledXY(this.win.Bounds().Center(), pixel.V(0.15, 0.15))
	matrix = matrix.Rotated(this.win.Bounds().Center(), this.angle)
	this.sprite.Draw(this.win, matrix)
}

func (this *hussein) RotateLeft(deltaTime float64) {
	this.angle += deltaTime * this.rotateSpeed
	this.Draw()
}

func (this *hussein) RotateRight(deltaTime float64) {
	this.angle -= deltaTime * this.rotateSpeed
	this.Draw()
}

func (this *hussein) ShootLaser() {
	this.lasers = append(this.lasers, NewLaser(this.resourceLoader, this.win, this.angle))
}

func (this *hussein) DrawLasers( /*Accepts some list of objects to check lasers against*/ ) {
	for i := 0; i < len(this.lasers); i++ {
		if this.lasers[i].Draw() {
			this.lasers[len(this.lasers)-1], this.lasers[i] = this.lasers[i], this.lasers[len(this.lasers)-1]
			this.lasers = this.lasers[:len(this.lasers)-1]
			i--
		} else {
			// Call checkCollisions for each laser
		}
	}
}
func (this *hussein) checkCollisions( /*Accepts some list of objects to check lasers against*/ ) {

}
