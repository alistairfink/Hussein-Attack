package entities

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
)

func generateEdgeSpawn(win *pixelgl.Window) (float64, float64) {
	var x, y float64
	edge := rand.Intn(4)
	bounds := win.Bounds()

	if edge == 0 {
		// Top
		y = bounds.Max.Y
		x = float64(rand.Intn(int(bounds.Max.X)))
	} else if edge == 1 {
		// Right
		x = bounds.Max.X
		y = float64(rand.Intn(int(bounds.Max.Y)))
	} else if edge == 2 {
		// Bottom
		y = -50
		x = float64(rand.Intn(int(bounds.Max.X)))
	} else {
		// Left
		x = -50
		y = float64(rand.Intn(int(bounds.Max.Y)))
	}

	return x, y
}

func calculateUnitSteps(start pixel.Vec, end pixel.Vec) (float64, float64) {
	deltaX := end.X - start.X
	deltaY := end.Y - start.Y
	total := abs(deltaX) + abs(deltaY)
	return deltaX / total, deltaY / total
}

func abs(val float64) float64 {
	if val < 0 {
		return -val
	}

	return val
}
