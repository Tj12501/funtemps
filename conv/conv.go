package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(farhenheit float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	return (farhenheit - 32) * (5 / 9)
}

// De andre konverteringsfunksjonene implementere her
// ...

func CelsiusToFarhenheit(celsius float64) float64 {
	return (celsius*1.8 + 32)
}

func CelsiusToKelvin(celsius float64) float64 {
	return (celsius + 273.15)
}

func KelvinToCelsius(kelvin float64) float64 {
	return (kelvin - 273.15)
}

func FarhenheitToKelvin(farhenheit float64) float64 {
	return ((farhenheit-32)*0.5555555555555556 + 273.15)
}

func KelvinToFarhenheit(kelvin float64) float64 {
	return ((kelvin-273.15)*1.8 + 32)
}
