package filters

import (
	"image"
	"image/color"
)

func Grayscale(img image.Image, method string) image.Image {
	imgBounds := img.Bounds()

	grayscaledImg := image.NewRGBA(image.Rect(0, 0, imgBounds.Max.X, imgBounds.Max.Y))
	grayscaledImgBounds := grayscaledImg.Bounds()

	width := grayscaledImgBounds.Max.X
	height := grayscaledImgBounds.Max.Y

	for x := grayscaledImgBounds.Min.X; x < width; x++ {
		for y := grayscaledImgBounds.Min.Y; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()

			var yComp uint8
			if method == "avg" {
				yComp = uint8((r + g + b) / 256 / 3)

			} else {
				yComp = uint8((0.299*float64(r) + 0.5870*float64(g) + 0.1140*float64(b)) / 256)
			}

			pixel := color.Gray{
				yComp,
			}

			grayscaledImg.Set(x, y, pixel)
		}
	}

	return grayscaledImg
}
