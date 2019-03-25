package main

import (
	"image"
	"image/png"
	"os"

	"github.com/shozawa/go_pl/ch03/hue"
)

func main() {
	const (
		width, height = 1024, 512
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, hue.Coloring(float64(x)/width))
		}
	}
	png.Encode(os.Stdout, img)
}
