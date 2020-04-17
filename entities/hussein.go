package entities

import (
	"github.com/alistairfink/Hussein-Attack/constants"
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
	laserCooldown  int
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
	obj.laserCooldown = constants.LaserCooldown

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
	if this.laserCooldown <= 0 {
		this.lasers = append(this.lasers, NewLaser(this.resourceLoader, this.win, this.angle))
		this.laserCooldown = constants.LaserCooldown
	}
}

func (this *hussein) DrawLasers(toiletPaperRolls *[]ToiletPaper) map[int]bool {
	toiletPaperResult := make(map[int]bool)
	for i := 0; i < len(this.lasers); i++ {
		if this.lasers[i].Draw() {
			this.removeLaser(i)
			i--
		} else {
			if this.checkCollisions(toiletPaperRolls, toiletPaperResult, i) {
				this.removeLaser(i)
				i--
			}
		}
	}

	if this.laserCooldown > 0 {
		this.laserCooldown--
	}

	return toiletPaperResult
}
func (this *hussein) checkCollisions(toiletPaperRolls *[]ToiletPaper, toiletPaperResult map[int]bool, currLaser int) bool {
	absPosX, absPosY := this.win.Bounds().Center().X+this.lasers[currLaser].posX, this.win.Bounds().Center().Y+this.lasers[currLaser].posY
	for i, roll := range *toiletPaperRolls {
		if abs(absPosX-roll.posX) < 10 && abs(absPosY-roll.posY) < 10 {
			toiletPaperResult[i] = true
			return true
		}
	}

	return false
}

func (this *hussein) removeLaser(i int) {
	this.lasers[len(this.lasers)-1], this.lasers[i] = this.lasers[i], this.lasers[len(this.lasers)-1]
	this.lasers = this.lasers[:len(this.lasers)-1]
}
