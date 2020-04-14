package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image"
	_ "image/png"
	"os"
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
	husseinSprite.Draw(win, pixel.IM.Moved((win.Bounds().Center())))

	for !win.Closed() {
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
