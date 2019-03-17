package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/shozawa/go_pl/ch02/lenconv"
	"github.com/shozawa/go_pl/ch02/tempconv"
	"github.com/shozawa/go_pl/ch02/weightconv"
)

var measure = flag.String("m", "temp", "measure")

func main() {
	flag.Parse()
	var values []string
	args := flag.Args()

	if len(args) < 1 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			values = append(values, input.Text())
		}
	} else {
		values = args
	}
	for _, value := range values {
		t, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "c: %v\n", err)
			os.Exit(1)
		}

		if *measure == "weight" {
			toWeight(t)
		} else if *measure == "len" {
			toLen(t)
		} else {
			toTemp(t)
		}
	}
}

func toTemp(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func toWeight(t float64) {
	k := weightconv.Kilogram(t)
	p := weightconv.Pound(t)
	fmt.Printf("%s = %s, %s = %s\n", k, weightconv.KtoP(k), p, weightconv.PtoK(p))
}

func toLen(t float64) {
	m := lenconv.Metre(t)
	f := lenconv.Feet(t)
	fmt.Printf("%s = %s, %s = %s\n", m, lenconv.MToF(m), f, lenconv.FToM(f))
}
