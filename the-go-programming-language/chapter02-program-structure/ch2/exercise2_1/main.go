// Exercise 2.1: Add types, constants, and functions to tempconv for processing temperatures
// in the Kelvin scale, where zero Kelvin is −273.15°C and a difference of 1K has the same
// magnitude as 1°C.
package main

import "fmt"

func main() {
	fmt.Println(CToK(-273.15))
	fmt.Println(KtoC(0))
}

func init() {
  fmt.Println("hello")
}

type Celsius float64

type Fahrenheit float64

type Kelvin float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (f Kelvin) String() string     { return fmt.Sprintf("%gK", f) }

const (
	AbsoulateZeroC Celsius = -273.15
	FreezingC      Celsius = 0
	BoilingC       Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KtoC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
