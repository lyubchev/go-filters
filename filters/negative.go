package filters

import (
	"image"
	"image/color"
)

// Negative is a function that takes an image and returns a color inverted version of it.
func Negative(img image.Image) image.Image {
	imgBounds := img.Bounds()

	nImg := image.NewRGBA(image.Rect(0, 0, imgBounds.Max.X, imgBounds.Max.Y))
	nImgBounds := nImg.Bounds()

	width := nImgBounds.Max.X
	height := nImgBounds.Max.Y

	for x := nImgBounds.Min.X; x < width; x++ {
		for y := nImgBounds.Min.Y; y < height; y++ {
			r, g, b, a := img.At(x, y).RGBA()

			pixel := color.RGBA{
				uint8(255 - (r / 256)),
				uint8(255 - (g / 256)),
				uint8(255 - (b / 256)),
				uint8(a),
			}

			nImg.Set(x, y, pixel)
		}
	}

	return nImg
}
