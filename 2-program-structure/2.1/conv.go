package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func KtoC(f Kelvin) Celsius { return Celsius(f + 273.15) }

func KtoF(f Kelvin) Fahrenheit { return Fahrenheit((f-273.15)*9/5 + 32) }

func FtoK(f Fahrenheit) Kelvin { return Kelvin(((f - 32) * 5 / 9) + 273.15) }

func CtoK(f Celsius) Kelvin { return Kelvin(f - 273.15) }
