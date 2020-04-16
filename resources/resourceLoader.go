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

func (this *ResourceLoader) loadImage(filePath string) (pixel.Picture, error) {
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

func (this *ResourceLoader) LoadHussein() pixel.Picture {
	husseinImageData, err := this.loadImage(husseinImageName)
	if err != nil {
		panic(err.Error())
	}

	return husseinImageData
}

func (this *ResourceLoader) LoadLaser() pixel.Picture {
	laserImageData, err := this.loadImage(laserImageName)
	if err != nil {
		panic(err.Error())
	}

	return laserImageData
}
