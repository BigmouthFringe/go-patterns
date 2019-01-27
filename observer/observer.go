package observer

// Observer Pattern defines one-to-many dependency between objects
// (to be specific - between Subject or Observable and Observers)
// so that when one object changes state, all its dependents are notified and updated automatically
type Observer interface {
	Update(temp, humidity, pressure float32)
}
type Observable interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

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
