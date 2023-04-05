package tempflag

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

type CelsiusFlagType struct{ Celsius }

// pflags package requires implementing the Value interface, which means implementing below String()/Type()/Set()
func (f *CelsiusFlagType) String() string {
	return f.Celsius.String()
}

func (f *CelsiusFlagType) Type() string {
	t, _ := fmt.Printf("%T", f)
	return t
}

func (f *CelsiusFlagType) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
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
