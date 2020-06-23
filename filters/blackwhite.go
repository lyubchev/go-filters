package filters

import (
	"image"
	"image/color"
)

// BlackWhite is a function that takes an image and a method and returns a black&whhite version of it
func BlackWhite(img image.Image) image.Image {
	imgBounds := img.Bounds()

	bwImg := image.NewRGBA(image.Rect(0, 0, imgBounds.Max.X, imgBounds.Max.Y))
	bwImgBounds := bwImg.Bounds()

	width := bwImgBounds.Max.X
	height := bwImgBounds.Max.Y

	for x := bwImgBounds.Min.X; x < width; x++ {
		for y := bwImgBounds.Min.Y; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()

			yComp := uint8((r + g + b) / 256 / 3)

			if yComp > 128 {
				yComp = 255
			} else {
				yComp = 0
			}

			pixel := color.Gray{
				yComp,
			}

			bwImg.Set(x, y, pixel)
		}
	}

	return bwImg
}
