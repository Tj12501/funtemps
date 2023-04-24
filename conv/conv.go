package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FahrenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFahrenheit
    ...
*/

// Konverterer Fahrenheit til Celsius
func FahrenheitToCelsius(fahrenheit float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	return (fahrenheit - 32) * (0.5556)
}

// De andre konverteringsfunksjonene implementere her
// ...

func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 1.8) + 32
}

func CelsiusToKelvin(celsius float64) float64 {
	return (celsius + 273.15)
}

func KelvinToCelsius(kelvin float64) float64 {
	return (kelvin - 273.15)
}

func FahrenheitToKelvin(fahrenheit float64) float64 {
	return ((fahrenheit-32)*0.556 + 273.15)
}

func KelvinToFahrenheit(kelvin float64) float64 {
	return ((kelvin-273.15)*1.8 + 32)
}
