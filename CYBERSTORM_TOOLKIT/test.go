package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"strconv"
)

func processPixels(file image.Image) []string {
	var binArray []string
	MaxX, MaxY := file.Bounds().Max.X, file.Bounds().Max.Y

	for x := 0; x < MaxX; x++ {
		for y := 0; y < MaxY; y++ {
			r, g, b, a := file.At(x, y).RGBA()
			r /= 257
			g /= 257
			b /= 257
			a /= 257
			//fmt.Printf("Pixel at (%d, %d): R=%s, G=%s, B=%s, A=%s\n", x, y, strconv.FormatInt(int64(r), 2), strconv.FormatInt(int64(g), 2), strconv.FormatInt(int64(b), 2), strconv.FormatInt(int64(a), 2))
			binArray = append(binArray, strconv.FormatInt(int64(r), 2), strconv.FormatInt(int64(g), 2), strconv.FormatInt(int64(b), 2), strconv.FormatInt(int64(a), 2))
		}
	}
	return binArray
}

func main() {
	file, err := os.Open("image.png")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	cat, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	//fmt.Println(cat)

	list := processPixels(cat)
	fmt.Println(list)
}
