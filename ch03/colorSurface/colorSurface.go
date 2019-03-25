package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	render(corner)
}

func render(f func(float64, float64) float64) {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := polygon(i+1, j, f)
			bx, by := polygon(i, j, f)
			cx, cy := polygon(i, j+1, f)
			dx, dy := polygon(i+1, j+1, f)
			r, g, b, a := coloring(i, j, f)
			fmt.Printf("<polygon points='%g, %g, %g, %g, %g, %g, %g, %g' fill='rgba(%d, %d, %d, %d)'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, r, g, b, a)
		}
	}
	fmt.Println("</svg>")
}

func coloring(i, j int, f func(float64, float64) float64) (uint8, uint8, uint8, uint8) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y) + 0.1

	if z > 0.8 {
		return 255, 0, 0, 255
	} else if z > 0.6 {
		return 255, 255, 0, 255
	} else if z > 0.4 {
		return 0, 255, 0, 255
	} else if z > 0.2 {
		return 0, 255, 255, 255
	} else {
		return 0, 0, 255, 255
	}
}

func polygon(i, j int, f func(float64, float64) float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	return projection(x, y, z)
}

func corner(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func projection(x, y, z float64) (float64, float64) {
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
