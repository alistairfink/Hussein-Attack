package resources

import (
	"github.com/faiface/pixel"
	"image"
	_ "image/png"
	"os"
)

type ResourceLoader struct{}

func NewResourceLoader() ResourceLoader {
	obj := ResourceLoader{}
	return obj
}

func (this *ResourceLoader) loadImage(filePath string) pixel.Picture {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		panic(err.Error())
	}

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err.Error())
	}

	return pixel.PictureDataFromImage(img)
}

func (this *ResourceLoader) LoadIcon() pixel.Picture {
	return this.loadImage(iconImageName)
}

func (this *ResourceLoader) LoadLogo() pixel.Picture {
	return this.loadImage(logoImageName)
}

func (this *ResourceLoader) LoadHussein() pixel.Picture {
	return this.loadImage(husseinImageName)
}

func (this *ResourceLoader) LoadLaser() pixel.Picture {
	return this.loadImage(laserImageName)
}

func (this *ResourceLoader) LoadToiletPaper() pixel.Picture {
	return this.loadImage(toiletPaperImageName)
}

func (this *ResourceLoader) LoadVirus() pixel.Picture {
	return this.loadImage(virusImageName)
}
