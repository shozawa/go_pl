package hue

// ref. https://qiita.com/masato_ka/items/c178a53c51364703d70b

import (
	"image/color"
	"math"
)

const (
	gain        = 10
	offsetx     = 0.2
	offsetGreen = 0.6
)

func Coloring(x float64) color.Color {
	x = x*2 - 1
	red := sigmoid(x, gain, -1*offsetx)
	blue := 1 - sigmoid(x, gain, offsetx)
	green := sigmoid(x, gain, offsetGreen) + (1 - sigmoid(x, gain, -1*offsetGreen)) - 1
	return color.RGBA{uint8(255 * red), uint8(255 * green), uint8(255 * blue), 255}
}

func sigmoid(x, gain, offsetx float64) float64 {
	return (math.Tanh((x+offsetx)*gain/2) + 1) / 2
}
