package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image"
	_ "image/png"
	"os"
	"time"
)

const GAME_TITLE = "2D Game"

func main() {
	pixelgl.Run(run)
}

func run() {
	config := pixelgl.WindowConfig{
		Title:  GAME_TITLE,
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(config)
	if err != nil {
		panic(err.Error())
	}

	hussein, err := loadImage("./resources/Hussein.png")
	if err != nil {
		panic(err.Error())
	}

	husseinSprite := pixel.NewSprite(hussein, hussein.Bounds())

	angle := 0.0

	last := time.Now()
	for !win.Closed() {
		win.Clear(colornames.Firebrick)

		deltaTime := time.Since(last).Seconds()
		last = time.Now()
		angle += 1.5 * deltaTime

		husseinMatrix := pixel.IM
		husseinMatrix = husseinMatrix.Moved(win.Bounds().Center())
		husseinMatrix = husseinMatrix.ScaledXY(win.Bounds().Center(), pixel.V(0.15, 0.15))
		husseinMatrix = husseinMatrix.Rotated(win.Bounds().Center(), angle)
		husseinSprite.Draw(win, husseinMatrix)

		win.Update()

	}
}

func loadImage(filePath string) (pixel.Picture, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}
