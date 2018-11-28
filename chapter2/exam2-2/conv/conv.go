package conv

//!+ temperature convert
// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k) }

//!- temperature convert

//!+ length convert
func MToK(m Meter) Kilometer { return Kilometer(m / 1000) }
func KToM(k Kilometer) Meter { return Meter(1000 * k) }

//!- length convert
func GToKG(g Grams) Kilograms  { return Kilograms(g / 1000) }
func KGToG(kg Kilograms) Grams { return Grams(1000 * kg) }

//!+ weight convert
