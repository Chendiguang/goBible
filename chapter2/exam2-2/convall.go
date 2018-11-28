//!+ convert temperature, length, weight
/*	Usage:
 *		go build goBible/chapter2/exam2-2/convall.go
 *		./convall -t 100.0 -l 10.0 -w 1
 *		Output:
 *			100°F = 37.77777777777778°C, 100°C = 212°F, 100K = 100°C
 *			10m = 0.01km, 10km = 10000m
 *			1g = 0.001kg, 1kg = 1000g
 */
//!-
package main

import (
	"flag"
	"fmt"
	"goBible/chapter2/exam2-2/conv"
)

var temp, length, weight float64

func init() {
	flag.Float64Var(&temp, "t", 100.0, "help mesage for temperature")
	flag.Float64Var(&length, "l", 100.0, "help mesage for length")
	flag.Float64Var(&weight, "w", 100.0, "help mesage for weight")
}

func main() {
	flag.Parse()
	//!+ temperature
	f := conv.Fahrenheit(temp)
	c := conv.Celsius(temp)
	k := conv.Kelvin(temp)
	fmt.Printf("%s = %s, %s = %s, %s = %s\n",
		f, conv.FToC(f), c, conv.CToF(c), k, conv.KToC(k))
	//!- temperature

	//!+ length
	meter := conv.Meter(length)
	kilomter := conv.Kilometer(length)
	fmt.Printf("%s = %s, %s = %s\n",
		meter, conv.MToK(meter), kilomter, conv.KToM(kilomter))
	//!- length

	//!+ weight
	grams := conv.Grams(weight)
	kilograms := conv.Kilograms(weight)
	fmt.Printf("%s = %s, %s = %s\n",
		grams, conv.GToKG(grams), kilograms, conv.KGToG(kilograms))
}
