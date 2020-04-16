package main

import (
	"github.com/alistairfink/Hussein-Attack/constants"
	"github.com/alistairfink/Hussein-Attack/entities"
	"github.com/alistairfink/Hussein-Attack/resources"
	"github.com/alistairfink/Hussein-Attack/state"
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

	// Startup
	resourceLoader := resources.NewResourceLoader()
	stateMachine := state.NewStateMachine()

	// Entities
	husseinEntity := entities.NewHussein(&resourceLoader, win)
	mainMenuEntity := entities.NewMainMenu(&resourceLoader, win)
	tp := entities.NewToiletPaper(&resourceLoader, win, 0.0)

	lastFrameTime := time.Now()
	for !win.Closed() {
		win.Clear(colornames.Black)

		deltaTime := time.Since(lastFrameTime).Seconds()
		lastFrameTime = time.Now()

		if stateMachine.IsMainMenu() {
			mainMenuEntity.Draw()

			if win.Pressed(pixelgl.KeySpace) {
				stateMachine.UpdateStateGameplay()
			}
		} else if stateMachine.IsGamePlay() {
			tp.Draw()
			if win.Pressed(pixelgl.KeyLeft) {
				husseinEntity.RotateLeft(deltaTime)
			} else if win.Pressed(pixelgl.KeyRight) {
				husseinEntity.RotateRight(deltaTime)
			} else {
				husseinEntity.Draw()
			}

			if win.Pressed(pixelgl.KeySpace) {
				husseinEntity.ShootLaser()
			}

			husseinEntity.DrawLasers()
		} else {
			panic("Error. Closing " + constants.GameTitle + ".")
		}

		win.Update()
	}
}
