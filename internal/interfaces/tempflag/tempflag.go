package tempflag

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

type CelsiusFlagType struct{ Celsius }

// pflags package requires implementing the Value interface, which means implementing below String()/Type()/Set()
// this is for the value the flag contains
func (f CelsiusFlagType) String() string {
	return f.Celsius.String()
}

func (f CelsiusFlagType) Type() string {
	t := fmt.Sprintf("%T", f)
	return t
}

func (f CelsiusFlagType) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		fmt.Println(f.Celsius)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		fmt.Println(f.Celsius)
		return nil
	case "K", "°K":
		f.Celsius = KtoC(Kelvin(value))
		fmt.Println(f.Celsius)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, shorthand string, value Celsius, usage string) *Celsius {
	f := CelsiusFlagType{value}
	flag.CommandLine.VarP(&f, name, shorthand, usage)
	return &f.Celsius
}

// copied from original tempconv exercise for extension
type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func KtoC(f Kelvin) Celsius     { return Celsius(f + 273.15) }
func KtoF(f Kelvin) Fahrenheit  { return Fahrenheit((f-273.15)*9/5 + 32) }
func FtoK(f Fahrenheit) Kelvin  { return Kelvin(((f - 32) * 5 / 9) + 273.15) }
func CtoK(f Celsius) Kelvin     { return Kelvin(f - 273.15) }

func CobraOutput() {
	fmt.Println(`
	Chapter 7 Exercise 7 has no executable program, heres a description instead:
	
	The exercise asked for why the help message contains degrees C unit, when the default value doesnt.
	This isnt strictly true for our effort.
	
	---------------------------------------------------------------------------------------------
	
	The reason for this is that the default value is a float64 however when it is read in via pflags
	it is wrapped in the CelsisusFlagType. This type is a wrapper around the underlying Celsius type 
	that satisfies a few interfaces, one of them is the stringer interface. 

	CelsiusFlagType -> Celsius -> Float64

	A float64 also satisfies this interface, and is what is used to display the 20.0 format.
	If we print a Celsius type we print with the units by calling its own String method.


	
	---------------------------------------------------------------------------------------------
			`)
}
