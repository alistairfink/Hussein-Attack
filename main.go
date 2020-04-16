package main

import (
	"github.com/alistairfink/Hussein-Attack/constants"
	"github.com/alistairfink/Hussein-Attack/entities"
	"github.com/alistairfink/Hussein-Attack/resources"
	"github.com/alistairfink/Hussein-Attack/state"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math/rand"
	"time"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	config := pixelgl.WindowConfig{
		Title:  constants.GameTitle,
		Bounds: pixel.R(0, 0, constants.GameWidth, constants.GameHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(config)
	if err != nil {
		panic(err.Error())
	}

	// Startup
	resourceLoader := resources.NewResourceLoader()
	stateMachine := state.NewStateMachine()
	rand.Seed(time.Now().UnixNano())

	// Entities
	husseinEntity := entities.NewHussein(&resourceLoader, win)
	mainMenuEntity := entities.NewMainMenu(&resourceLoader, win)
	tp1 := entities.NewToiletPaper(&resourceLoader, win)
	tp2 := entities.NewToiletPaper(&resourceLoader, win)
	tp3 := entities.NewToiletPaper(&resourceLoader, win)
	tp4 := entities.NewToiletPaper(&resourceLoader, win)
	tp5 := entities.NewToiletPaper(&resourceLoader, win)
	tp6 := entities.NewToiletPaper(&resourceLoader, win)
	tp7 := entities.NewToiletPaper(&resourceLoader, win)
	tp8 := entities.NewToiletPaper(&resourceLoader, win)
	tp9 := entities.NewToiletPaper(&resourceLoader, win)
	tp10 := entities.NewToiletPaper(&resourceLoader, win)
	tp11 := entities.NewToiletPaper(&resourceLoader, win)
	tp12 := entities.NewToiletPaper(&resourceLoader, win)
	tp13 := entities.NewToiletPaper(&resourceLoader, win)
	tp14 := entities.NewToiletPaper(&resourceLoader, win)
	tp15 := entities.NewToiletPaper(&resourceLoader, win)

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
			tp1.Draw()
			tp2.Draw()
			tp3.Draw()
			tp4.Draw()
			tp5.Draw()
			tp6.Draw()
			tp7.Draw()
			tp8.Draw()
			tp9.Draw()
			tp10.Draw()
			tp11.Draw()
			tp12.Draw()
			tp13.Draw()
			tp14.Draw()
			tp15.Draw()

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
