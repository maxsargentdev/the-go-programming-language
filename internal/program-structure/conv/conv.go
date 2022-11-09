package conv

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Fahrenheit float64
type Celsius float64
type Feet float64
type Meter float64
type Pound float64
type Kilo float64

const (
	PoundPerKilo = 2.205
	FeetPerMeter = 3.28084
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (f Feet) String() string       { return fmt.Sprintf("%gFt", f) }
func (m Meter) String() string      { return fmt.Sprintf("%gM", m) }
func (lbs Pound) String() string    { return fmt.Sprintf("%glbs", lbs) }
func (kgs Kilo) String() string     { return fmt.Sprintf("%gkgs", kgs) }

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func FttoM(f Feet) Meter        { return Meter(f / FeetPerMeter) }
func MtoFt(m Meter) Feet        { return Feet(m * FeetPerMeter) }
func PtoKgs(lbs Pound) Kilo     { return Kilo(lbs / PoundPerKilo) }
func KgstoP(kgs Kilo) Pound     { return Pound(kgs * PoundPerKilo) }

func Convert(measures []string) {

	if len(measures) <= 1 {
		// read from stfout here
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			fmt.Println(parseInput(input.Text()))
		}
	} else {
		// read from command line args here
		for _, arg := range measures {
			fmt.Println(parseInput(arg))
		}
	}
}

func parseInput(s string) string {
	if HasSuffix(s, "°C") {
		celsiusString := strings.TrimSuffix(s, "°C")
		celsiusFloat, _ := strconv.ParseFloat(celsiusString, 64)
		fahrenheitValue := CtoF(Celsius(celsiusFloat))
		return fahrenheitValue.String()
	}
	if HasSuffix(s, "°F") {
		fahrenheitString := strings.TrimSuffix(s, "°F")
		fahrenheitFloat, _ := strconv.ParseFloat(fahrenheitString, 64)
		celsiusValue := FtoC(Fahrenheit(fahrenheitFloat))
		return celsiusValue.String()
	}
	if HasSuffix(s, "Ft") {
		feetString := strings.TrimSuffix(s, "Ft")
		feetFloat, _ := strconv.ParseFloat(feetString, 64)
		meterValue := FttoM(Feet(feetFloat))
		return meterValue.String()
	}
	if HasSuffix(s, "M") {
		meterString := strings.TrimSuffix(s, "M")
		meterFloat, _ := strconv.ParseFloat(meterString, 64)
		feetValue := MtoFt(Meter(meterFloat))
		return feetValue.String()
	}
	if HasSuffix(s, "lbs") {
		lbsString := strings.TrimSuffix(s, "lbs")
		lbsFloat, _ := strconv.ParseFloat(lbsString, 64)
		kgsValue := PtoKgs(Pound(lbsFloat))
		return kgsValue.String()
	}
	if HasSuffix(s, "kgs") {
		kgsString := strings.TrimSuffix(s, "kgs")
		kgsFloat, _ := strconv.ParseFloat(kgsString, 64)
		lbsValue := KgstoP(Kilo(kgsFloat))
		return lbsValue.String()
	}
	return "I cannot conver this unit"
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
