// Package weather tells you the weather forecast.
package weather

// CurrentCondition is a string that represents the current condition.
var CurrentCondition string
// CurrentLocation is a string that represents the city you want the
// forecast for.
var CurrentLocation string

// Forecast is a function that takes in the CurrentLocation and
// CurrentCondition and returns a string that represents the forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
