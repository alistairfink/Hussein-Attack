package main

import (
	"github.com/alistairfink/2D-Game-Fun/constants"
	"github.com/alistairfink/2D-Game-Fun/entities"
	"github.com/alistairfink/2D-Game-Fun/resources"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	config := pixelgl.WindowConfig{
		Title:  constants.GameTitle,
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(config)
	if err != nil {
		panic(err.Error())
	}

	resourceLoader := resources.NewResourceLoader()
	husseinEntity := entities.NewHussein(&resourceLoader, win)

	lastFrameTime := time.Now()
	for !win.Closed() {
		win.Clear(colornames.Firebrick)

		deltaTime := time.Since(lastFrameTime).Seconds()
		lastFrameTime = time.Now()

		husseinEntity.Draw(deltaTime)

		win.Update()
	}
}
