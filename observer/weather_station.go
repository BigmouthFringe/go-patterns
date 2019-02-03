package observer

import "fmt"

type WeatherData struct {
	observers Observers
	temperature float32
	humidity float32
	pressure float32
}

func (WeatherData) New() *WeatherData {
	var ws = new(WeatherData)
	return ws
}

func (ws *WeatherData) RegisterObserver(observer Observer) {
	ws.observers = append(ws.observers, observer)
}
func (ws *WeatherData) RemoveObserver(observer Observer) {
	var pos = ws.observers.indexOf(observer)
	if pos >= 0 {
		ws.observers.remove(pos)
	}
}
func (ws *WeatherData) NotifyObservers() {
	for _, observer := range ws.observers {
		observer.Update(ws.temperature, ws.humidity, ws.pressure)
	}
}

func (ws *WeatherData) SetMeasurements(temp, humid, pres float32) {
	ws.temperature = temp
	ws.humidity = humid
	ws.pressure = pres
	ws.MeasurementsChanged()
}
func (ws *WeatherData) MeasurementsChanged() {
	ws.NotifyObservers()
}

func (ws *WeatherData) GetTemperature() float32 {
	return ws.temperature
}
func (ws *WeatherData) GetHumidity() float32 {
	return ws.humidity
}
func (ws *WeatherData) GetPressure() float32 {
	return ws.pressure
}

type Observers []Observer
func (slice Observers) indexOf(observer Observer) int {
	for pos, o := range slice {
		if o == observer {
			return pos
		}
	}
	return -1
}
func (slice Observers) remove(pos int) Observers {
	return append(slice[:pos], slice[pos+1:]...)
}

// The result that we have is that after we change our WeatherData state
// it automatically notifies all our displays and they handle it (e.g. print out the messages)
func UpdateWeatherData() {
	var weatherData = WeatherData{}.New()
	CurrentConditionDisplay{}.New(weatherData)
	ForecastDisplay{}.New(weatherData)

	// Simulation of new weather measurements
	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)
}

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