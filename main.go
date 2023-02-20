package main

import (
	"flag"
	"fmt"
)

var fahr float64
var kelvin float64
var celsius float64
var out string
var funfacts string

func init() {
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i fahrenheit")
	flag.Float64Var(&kelvin, "K", 0.0, "Temperatur i kelvin")
	flag.Float64Var(&celsius, "C", 0.0, "Temperature i celsius")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
}

func main() {
	flag.Parse()

	fmt.Println(fahr, out, funfacts)
	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())
	fmt.Println(isFlagPassed("out"))

	if out == "C" && isFlagPassed("F") {
		fmt.Println("0°F er -17.78°C")
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
