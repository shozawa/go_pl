package main

import (
	"errors"
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
	// mogul
	render(func(x, y float64) float64 { return math.Sin(x) * math.Cos(y) / 4 })
}

func render(f func(float64, float64) float64) {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var err error
			ax, ay, err := polygon(i+1, j, f)
			bx, by, err := polygon(i, j, f)
			cx, cy, err := polygon(i, j+1, f)
			dx, dy, err := polygon(i+1, j+1, f)

			if err != nil {
				continue
			}

			fmt.Printf("<polygon points='%g, %g, %g, %g, %g, %g, %g, %g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func polygon(i, j int, f func(float64, float64) float64) (float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	if z < math.Inf(1) && z > math.Inf(-1) {
		sx, sy := projection(x, y, z)
		return sx, sy, nil
	} else {
		return 0, 0, errors.New("should not infinity")
	}
}

func projection(x, y, z float64) (float64, float64) {
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
