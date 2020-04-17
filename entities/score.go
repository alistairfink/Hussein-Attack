package entities

import (
	"fmt"
	"github.com/alistairfink/Hussein-Attack/constants"
	"github.com/alistairfink/Hussein-Attack/resources"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type score struct {
	win   *pixelgl.Window
	atlas *text.Atlas
	score int
}

func NewScore(resourceLoader *resources.ResourceLoader, win *pixelgl.Window) score {
	obj := score{}
	obj.win = win
	obj.atlas = text.NewAtlas(
		basicfont.Face7x13,
		text.ASCII,
	)
	obj.score = 0

	return obj
}

func (this *score) Draw() {
	scoreText := text.New(pixel.V(0, 0), this.atlas)
	fmt.Fprintln(scoreText, constants.ScoreText, this.score)
	scoreTextPos := pixel.IM.Moved(pixel.V(20, this.win.Bounds().Max.Y-scoreText.Bounds().Max.Y-20))

	scoreText.Draw(this.win, scoreTextPos.Scaled(pixel.V(20, this.win.Bounds().Max.Y-scoreText.Bounds().Max.Y-20), 1.4))
}

func (this *score) IncrementScore(increment int) {
	this.score += increment
}

func (this *score) Score() int {
	return this.score
}

func (this *score) Reset() {
	this.score = 0
}
