package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

func main() {
	img, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer img.Close()

	filename := path.Base(os.Args[1])
	filename = strings.Trim(filename, path.Ext(filename))

	imgType := path.Ext(os.Args[1])

	var loadedImg image.Image
	if imgType == ".jpeg" || imgType == ".jpg" {
		loadedImg, err = jpeg.Decode(img)
	} else {
		loadedImg, err = png.Decode(img)
	}

	fmt.Println(filename)
	if err != nil {
		panic(err)
	}

	grayscaledImg := grayscale(loadedImg)
	grayscaledImgFile, err := os.Create(fmt.Sprintf("%s-grayscaled.png", filename))
	if err != nil {
		panic(err)
	}

	err = png.Encode(grayscaledImgFile, grayscaledImg)

	if err != nil {
		panic(err)
	}

}

func grayscale(img image.Image) image.Image {
	imgBounds := img.Bounds()

	grayscaledImg := image.NewRGBA(image.Rect(0, 0, imgBounds.Max.X, imgBounds.Max.Y))
	grayscaledImgBounds := grayscaledImg.Bounds()

	width := grayscaledImgBounds.Max.X
	height := grayscaledImgBounds.Max.Y

	for x := grayscaledImgBounds.Min.X; x < width; x++ {
		for y := grayscaledImgBounds.Min.Y; y < height; y++ {
			r, g, b, a := img.At(x, y).RGBA()

			grey := uint8(0.2989*float64(r) + 0.5870*float64(g) + 0.1140*float64(b))
			alpha := uint8(a)

			rgba := color.RGBA{
				grey,
				grey,
				grey,
				alpha,
			}

			grayscaledImg.Set(x, y, rgba)
		}
	}

	return grayscaledImg
}
