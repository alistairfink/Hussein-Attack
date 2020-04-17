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
	var counter uint64 = 0
	viruseSpawnRate := constants.InitialViruseSpawnRate
	println(0, viruseSpawnRate)

	// Entities
	mainMenuEntity := entities.NewMainMenu(&resourceLoader, win)
	scoreEntity := entities.NewScore(&resourceLoader, win)
	husseinEntity := entities.NewHussein(&resourceLoader, win)
	toiletPaperEntities := []entities.ToiletPaper{}
	virusEntities := []entities.Virus{}

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
			// Virus Ramp Up
			counter++
			if viruseSpawnRate > constants.VirusSpawnRateMin && counter%constants.VirusSpawnRateRampUp == 0 {
				viruseSpawnRate--
			}

			// Score
			scoreEntity.Draw()

			// Toilet Paper
			toiletPaperRandomNum := rand.Intn(constants.ToiletPaperSpawnRate)
			if toiletPaperRandomNum == 0 {
				toiletPaperEntities = append(toiletPaperEntities, entities.NewToiletPaper(&resourceLoader, win))
			}

			for i := 0; i < len(toiletPaperEntities); i++ {
				if toiletPaperEntities[i].Draw() {
					toiletPaperEntities[i], toiletPaperEntities[len(toiletPaperEntities)-1] = toiletPaperEntities[len(toiletPaperEntities)-1], toiletPaperEntities[i]
					toiletPaperEntities = toiletPaperEntities[:len(toiletPaperEntities)-1]
					scoreEntity.IncrementScore(constants.ToiletPaperScore)
					i--
				}
			}

			// Viruses
			virusRnadomNum := rand.Intn(viruseSpawnRate)
			if virusRnadomNum == 0 {
				virusEntities = append(virusEntities, entities.NewVirus(&resourceLoader, win))
			}

			for i := 0; i < len(virusEntities); i++ {
				if virusEntities[i].Draw() {
					// Gameover
					// Change State here
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

			shotToiletPaper, shotViruses := husseinEntity.DrawLasers(&toiletPaperEntities, &virusEntities)
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

			if len(shotViruses) > 0 {
				newVirusEntities := make([]entities.Virus, len(virusEntities)-len(shotViruses))
				i := 0
				for j := 0; j < len(virusEntities); j++ {
					if !shotViruses[j] {
						newVirusEntities[i] = virusEntities[j]
						i++
					}
				}

				virusEntities = newVirusEntities
			}
		} else {
			panic("Error. Closing " + constants.GameTitle + ".")
		}

		win.Update()
	}
}
