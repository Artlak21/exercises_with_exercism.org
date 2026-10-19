// Package weather stores current weather data and generates a forecast.
package weather

var (
    // CurrentCondition — current weather condition.
	CurrentCondition string
    // CurrentLocation — current location.
	CurrentLocation  string
)

// Forecast updates the current data and returns a row with the city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
