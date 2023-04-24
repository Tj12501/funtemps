package main

import (
	"flag"
	"fmt"
	"github.com/Tj12501/funtemps/conv"
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
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - fahrenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
}

func main() {
	flag.Parse()

	if out == "C" && isFlagPassed("F") {
		fmt.Printf("%v°F er = %.2f°C", fahr, conv.FahrenheitToCelsius(fahr))
	}

	if out == "C" && isFlagPassed("K") {
		fmt.Printf("%v°K er = %2.f°C", kelvin, conv.KelvinToCelsius(kelvin))
	}

	if out == "F" && isFlagPassed("C") {
		fmt.Printf("%v°C er = %2.f°F", celsius, conv.CelsiusToFahrenheit(celsius))
	}

	if out == "F" && isFlagPassed("K") {
		fmt.Printf("%v°K er = %2.f°F", kelvin, conv.KelvinToFahrenheit(kelvin))
	}

	if out == "K" && isFlagPassed("C") {
		fmt.Printf("%v°C er = %2.f°K", celsius, conv.CelsiusToKelvin(celsius))
	}

	if out == "K" && isFlagPassed("F") {
		fmt.Printf("%v°F er = %.2f°K", fahr, conv.FahrenheitToKelvin(fahr))
	}
}

/*
Kun en kopi fra dokumentasjon til fmt pakken
%f     default width, default precision
%9f    width 9, default precision
%.2f   default width, precision 2
%9.2f  width 9, precision 2
%9.f   width 9, precision 0
*/
func FormatBigInteger(number float64) string {

	s := fmt.Sprintf("%.f", number)

	var sc string
	var scc string
	var neg bool

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == "-" {
			neg = true
			continue
		}
		if len(sc) < 3 {
			sc = string(s[i]) + sc
			//fmt.Println("IF: i", i, "sc = ", sc, "len(sc) = ", len(sc))
		} else {
			//fmt.Println("ELSE START: scc = ", scc, "sc = ", sc)
			scc = sc + " " + scc
			sc = string(s[i])
			//fmt.Println("ELSE END: scc = ", scc, "sc = ", sc)
		}
	}
	//fmt.Println("len(sc) = ", len(sc), "sc = ", sc, "scc = ", scc)
	if neg {
		sc = "-" + sc
	}
	if len(scc) > 0 {
		return sc + " " + scc[:len(scc)-1]
	} else {
		return sc
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
