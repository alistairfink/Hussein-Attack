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

type mainMenu struct {
	image        pixel.Picture
	sprite       *pixel.Sprite
	win          *pixelgl.Window
	atlas        *text.Atlas
	startText    *text.Text
	startTextPos pixel.Matrix
}

func NewMainMenu(resourceLoader *resources.ResourceLoader, win *pixelgl.Window) mainMenu {
	obj := mainMenu{}
	obj.win = win
	obj.atlas = text.NewAtlas(
		basicfont.Face7x13,
		text.ASCII,
	)
	obj.startText = text.New(pixel.V(0, 0), obj.atlas)
	fmt.Fprintln(obj.startText, constants.StartGameText)
	obj.startTextPos = pixel.IM.Moved(pixel.V(win.Bounds().Center().Sub(obj.startText.Bounds().Center()).X, 150))

	return obj
}

func (this *mainMenu) Draw() {
	this.startText.Draw(this.win, this.startTextPos)
}
