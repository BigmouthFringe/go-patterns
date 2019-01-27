package observer

import "fmt"

type Display interface {
	Display()
}

type CurrentConditionDisplay struct {
	temperature float32
	humidity    float32
	weatherData *WeatherData
}
func (CurrentConditionDisplay) New(weatherData *WeatherData) *CurrentConditionDisplay {
	var display = new(CurrentConditionDisplay)
	display.weatherData = weatherData
	weatherData.RegisterObserver(display)
	return display
}
func (ccd *CurrentConditionDisplay) Update(temp, humid, pres float32) {
	ccd.temperature = temp
	ccd.humidity = humid
	ccd.Display()
}
func (ccd *CurrentConditionDisplay) Display() {
	var message = fmt.Sprintf("Current conditions: %vF degrees and %v%% humidity", ccd.temperature, ccd.humidity)
	fmt.Println(message)
}

type ForecastDisplay struct {
	currentPressure float32
	lastPressure float32
	weatherData *WeatherData
}
func (ForecastDisplay) New(weatherData *WeatherData) *ForecastDisplay {
	var display = new(ForecastDisplay)
	display.currentPressure = 29.92
	display.weatherData = weatherData
	weatherData.RegisterObserver(display)
	return display
}
func (fd *ForecastDisplay) Update(temp, humid, pres float32) {
	fd.lastPressure = fd.currentPressure
	fd.currentPressure = fd.weatherData.GetPressure()
	fd.Display()
}
func (fd *ForecastDisplay) Display() {
	var message = "More of the same"
	if fd.currentPressure > fd.lastPressure {
		message = "Improving weather on the way"
	}
	if fd.currentPressure < fd.lastPressure {
		message = "Watch out for cooler, rainy weather"
	}
	fmt.Println("Forecast:", message)
}