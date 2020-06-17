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
	method := os.Args[2]

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

	grayscaledImg := grayscale(loadedImg, method)
	grayscaledImgFile, err := os.Create(fmt.Sprintf("%s-grayscaled-%s%s", filename, method, imgType))
	if err != nil {
		panic(err)
	}
	defer grayscaledImgFile.Close()

	err = jpeg.Encode(grayscaledImgFile, grayscaledImg, nil)

	if err != nil {
		panic(err)
	}

}

func grayscale(img image.Image, method string) image.Image {
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
