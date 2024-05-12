package weather

type Subject interface {
	RegistObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type WeatherData struct {
	observers []Observer
	temperture float32
	humidity float32
	pressure float32
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (data *WeatherData) RegistObserver(o Observer) {
	data.observers = append(data.observers, o)
}

func (data *WeatherData) RemoveObserver(o Observer) {
	var observers []Observer

	for _, observer := range data.observers {
		if observer != o {
			observers = append(observers, observer)
		}
	}

	data.observers = observers
}

func (data *WeatherData) NotifyObservers() {
	for _, observer := range data.observers {
		observer.update(data.temperture, data.humidity, data.pressure)
	}
}

func (data *WeatherData) MeasurementChanged() {
	data.NotifyObservers()
}

func (data *WeatherData) SetMessurement(temp float32, humidity float32, pressure float32) {
	data.temperture = temp
	data.humidity = humidity
	data.pressure = pressure
	data.MeasurementChanged()
}