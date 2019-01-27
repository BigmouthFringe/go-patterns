package observer

func UpdateWeatherData() {
	var weatherData = WeatherData{}.New()
	CurrentConditionDisplay{}.New(weatherData)
	ForecastDisplay{}.New(weatherData)

	// Simulation of new weather measurements
	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)
}