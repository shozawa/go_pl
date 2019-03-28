package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

type Options struct {
	width  int
	height int
}

func (o *Options) ParseQuery(q map[string]string) {
	if q["width"] != "" {
		o.width, _ = strconv.Atoi(q["width"])
	} else {
		o.width = width
	}
	if q["height"] != "" {
		o.height, _ = strconv.Atoi(q["height"])
	} else {
		o.height = height
	}
}

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", hander)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func params(r *http.Request) map[string]string {
	q := make(map[string]string)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		q[k] = v[0]
	}
	return q
}

func hander(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	var o Options
	o.ParseQuery(params(r))
	log.Println(o)
	render(w, o)
}

func render(out io.Writer, options Options) {
	width, height := options.width, options.height
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := polygon(i+1, j, width, height)
			bx, by := polygon(i, j, width, height)
			cx, cy := polygon(i, j+1, width, height)
			dx, dy := polygon(i+1, j+1, width, height)
			fmt.Fprintf(out, "<polygon points='%g, %g, %g, %g, %g, %g, %g, %g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprint(out, "</svg>")
}

func polygon(i, j, width, height int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// corner
	r := math.Hypot(x, y)
	z := math.Sin(r) / r
	// projection
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
