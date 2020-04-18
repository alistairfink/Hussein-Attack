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
	win          *pixelgl.Window
	atlas        *text.Atlas
	logoImage    pixel.Picture
	logoSprite   *pixel.Sprite
	logoPos      pixel.Matrix
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

	// Logo Image
	obj.logoImage = (*resourceLoader).LoadLogo()
	obj.logoSprite = pixel.NewSprite(obj.logoImage, obj.logoImage.Bounds())
	obj.logoPos = pixel.IM.Moved(pixel.V(win.Bounds().Center().X, win.Bounds().Max.Y-150))

	// Menu Text
	obj.menuText = text.New(pixel.V(0, 0), obj.atlas)
	for _, line := range constants.MenuText {
		obj.menuText.Dot.X -= obj.menuText.BoundsOf(line).W() / 2
		fmt.Fprintln(obj.menuText, line)
	}

	obj.menuTextPos = pixel.IM.Moved(win.Bounds().Center().Sub(obj.menuText.Bounds().Center()))
	obj.menuTextPos = obj.menuTextPos.Scaled(win.Bounds().Center(), 1.4)

	// Start Text
	obj.startText = text.New(pixel.V(0, 0), obj.atlas)
	fmt.Fprintln(obj.startText, constants.StartGameText)
	obj.startTextPos = pixel.IM.Moved(pixel.V(win.Bounds().Center().Sub(obj.startText.Bounds().Center()).X, 200))
	obj.startTextPos = obj.startTextPos.Scaled(win.Bounds().Center(), 1.4)

	return obj
}

func (menu *mainMenu) Draw() {
	menu.logoSprite.Draw(menu.win, menu.logoPos)
	menu.menuText.Draw(menu.win, menu.menuTextPos)
	menu.startText.Draw(menu.win, menu.startTextPos)
}
