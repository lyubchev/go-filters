package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
	"github.com/impzero/go-grayscale/filters"
)

func main() {
	img, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer img.Close()

	filename := path.Base(os.Args[1])
	filename = strings.TrimSuffix(filename, path.Ext(filename))

	imgType := path.Ext(os.Args[1])
	method := os.Args[2]
	if !(method == "coeff" || method == "avg") {
		panic("go-grayscale: method must be of type coeff or avg")
	}

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

	grayscaledImg := filters.Grayscale(loadedImg, method)
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
