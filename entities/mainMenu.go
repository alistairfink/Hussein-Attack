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
	menuText     *text.Text
	menuTextPos  pixel.Matrix
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

	// Menu Text
	obj.menuText = text.New(pixel.V(0, 0), obj.atlas)
	for _, line := range constants.MenuText {
		obj.menuText.Dot.X -= obj.menuText.BoundsOf(line).W() / 2
		fmt.Fprintln(obj.menuText, line)
	}

	obj.menuTextPos = pixel.IM.Moved(win.Bounds().Center().Sub(obj.menuText.Bounds().Center()))

	// Start Text
	obj.startText = text.New(pixel.V(0, 0), obj.atlas)
	fmt.Fprintln(obj.startText, constants.StartGameText)
	obj.startTextPos = pixel.IM.Moved(pixel.V(win.Bounds().Center().Sub(obj.startText.Bounds().Center()).X, 150))

	return obj
}

func (this *mainMenu) Draw() {
	this.menuText.Draw(this.win, this.menuTextPos.Scaled(this.win.Bounds().Center(), 1.4))
	// this.menuText.Draw(this.win, this.menuTextPos)
	this.startText.Draw(this.win, this.startTextPos.Scaled(this.win.Bounds().Center(), 1.4))
}
