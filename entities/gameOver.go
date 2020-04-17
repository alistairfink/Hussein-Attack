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

type gameOver struct {
	win             *pixelgl.Window
	atlas           *text.Atlas
	gameOverText    *text.Text
	gameOverTextPos pixel.Matrix
	restartText     *text.Text
	restartTextPos  pixel.Matrix
	score           int
}

func NewGameOver(resourceLoader *resources.ResourceLoader, win *pixelgl.Window) gameOver {
	obj := gameOver{}
	obj.win = win
	obj.atlas = text.NewAtlas(
		basicfont.Face7x13,
		text.ASCII,
	)
	obj.score = -1

	// Game Over Text
	obj.gameOverText = text.New(pixel.V(0, 0), obj.atlas)
	fmt.Fprintln(obj.gameOverText, constants.GameOverText)
	obj.gameOverTextPos = pixel.IM.Moved(pixel.V(win.Bounds().Center().Sub(obj.gameOverText.Bounds().Center()).X, win.Bounds().Max.Y-350))
	obj.gameOverTextPos = obj.gameOverTextPos.Scaled(win.Bounds().Center(), 6)

	// Restart Text
	obj.restartText = text.New(pixel.V(0, 0), obj.atlas)
	fmt.Fprintln(obj.restartText, constants.RestartGameText)
	obj.restartTextPos = pixel.IM.Moved(pixel.V(win.Bounds().Center().Sub(obj.restartText.Bounds().Center()).X, 200))
	obj.restartTextPos = obj.restartTextPos.Scaled(win.Bounds().Center(), 1.4)

	return obj
}

func (this *gameOver) Draw() {
	scoreText := text.New(pixel.V(0, 0), this.atlas)
	fmt.Fprintln(scoreText, constants.ScoreText, this.score)
	scoreTextPos := pixel.IM.Moved(pixel.V(this.win.Bounds().Center().Sub(scoreText.Bounds().Center()).X, this.win.Bounds().Center().Y))
	scoreTextPos = scoreTextPos.Scaled(this.win.Bounds().Center(), 3.5)

	this.gameOverText.Draw(this.win, this.gameOverTextPos)
	scoreText.Draw(this.win, scoreTextPos)
	this.restartText.Draw(this.win, this.restartTextPos)
}

func (this *gameOver) ScoreSet() bool {
	return this.score != -1
}

func (this *gameOver) SetScore(score int) {
	this.score = score
}

func (this *gameOver) Reset() {
	this.score = -1
}
