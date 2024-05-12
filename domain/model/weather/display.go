package weather

import (
	"fmt"
)

type Observer interface {
	update(temp float32, humidity float32, pressure float32)
}

type DisplayElement interface {
	display()
}

type CurrentConditionsDisplay struct {
	temperture float32
	humidity float32
	pressure float32
}

func NewCurrentConditionsDisplay() *CurrentConditionsDisplay {
	return &CurrentConditionsDisplay{}
}

func (d *CurrentConditionsDisplay) display() {
	fmt.Printf("CurrentConditionsDisplay %v, %v, %v\n", d.temperture, d.humidity, d.pressure)
}

func (d *CurrentConditionsDisplay) update(temp float32, humidity float32, pressure float32) {
	d.temperture = temp
	d.humidity = humidity
	d.pressure = pressure
	d.display()
}


