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
	mainMenuEntity := entities.NewMainMenu(&resourceLoader, win)
	scoreEntity := entities.NewScore(&resourceLoader, win)
	husseinEntity := entities.NewHussein(&resourceLoader, win)
	toiletPaperEntities := []entities.ToiletPaper{
		entities.NewToiletPaper(&resourceLoader, win),
		// entities.NewToiletPaper(&resourceLoader, win),
		// entities.NewToiletPaper(&resourceLoader, win),
		// entities.NewToiletPaper(&resourceLoader, win),
		// entities.NewToiletPaper(&resourceLoader, win),
		// entities.NewToiletPaper(&resourceLoader, win),
		// entities.NewToiletPaper(&resourceLoader, win),
		// entities.NewToiletPaper(&resourceLoader, win),
		// entities.NewToiletPaper(&resourceLoader, win),
		// entities.NewToiletPaper(&resourceLoader, win),
		// entities.NewToiletPaper(&resourceLoader, win),
		// entities.NewToiletPaper(&resourceLoader, win),
		// entities.NewToiletPaper(&resourceLoader, win),
		// entities.NewToiletPaper(&resourceLoader, win),
		// entities.NewToiletPaper(&resourceLoader, win),
	}

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
			// Score
			scoreEntity.Draw()

			// Toilet Paper
			for i := 0; i < len(toiletPaperEntities); i++ {
				if toiletPaperEntities[i].Draw() {
					toiletPaperEntities[i], toiletPaperEntities[len(toiletPaperEntities)-1] = toiletPaperEntities[len(toiletPaperEntities)-1], toiletPaperEntities[i]
					toiletPaperEntities = toiletPaperEntities[:len(toiletPaperEntities)-1]
					scoreEntity.IncrementScore(constants.ToiletPaperScore)
					i--
				}
			}

			// Hussein and Lasers
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

			shotToiletPaper := husseinEntity.DrawLasers(&toiletPaperEntities)
			if len(shotToiletPaper) > 0 {
				newToiletPaperEntities := make([]entities.ToiletPaper, len(toiletPaperEntities)-len(shotToiletPaper))
				i := 0
				for j := 0; j < len(toiletPaperEntities); j++ {
					if !shotToiletPaper[j] {
						newToiletPaperEntities[i] = toiletPaperEntities[j]
						i++
					}
				}

				toiletPaperEntities = newToiletPaperEntities
			}
		} else {
			panic("Error. Closing " + constants.GameTitle + ".")
		}

		win.Update()
	}
}
