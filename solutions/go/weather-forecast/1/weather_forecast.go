// Package weather provides tools to forecast weather condition
// for the location defined.
package weather

var (
    // CurrentCondition is a parameter used for saving the condition of the weather predicted.
	CurrentCondition string
    // CurrentLocation is a parameter used for saving the location of the weather prediction.
	CurrentLocation  string
)

// Forecast function returns a string defining for a city, the corresponding condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
