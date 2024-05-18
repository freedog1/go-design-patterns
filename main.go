package main

import (
	"design-pattern/domain/model/adventure"
	"design-pattern/domain/model/duck"
	"design-pattern/domain/model/starbuz"
	"design-pattern/domain/model/weather"
	"fmt"
)

func main() {
	// simulateDuck()
	// simulateKing()
	// simulateWeather()
	simulateStarbuz()
}

func simulateDuck() {
	miniDuck := duck.MiniDuckSimulator()
	miniDuck.Display()
	miniDuck.Duck.PerformQuack()
	miniDuck.Duck.PerformFly()

	modelDuck := duck.ModelDuckSimulator()
	modelDuck.Display()
	modelDuck.Duck.SetFlyBehavior(&duck.FlyRocketPowered{})
	modelDuck.Duck.PerformFly()
}


func simulateKing() {
	king := adventure.NewKing()
	king.Character.UseWeapon()
	fmt.Println("武器を変えて攻撃")
	king.Character.SetWeapon(&adventure.AxeBehavior{})
	king.Character.UseWeapon()
}

func simulateWeather() {
	weatherData := weather.NewWeatherData()
	weatherData.RegistObserver(weather.NewCurrentConditionsDisplay())
	weatherData.SetMessurement(10,20,30)
	weatherData.SetMessurement(100,200,300)
}

func simulateStarbuz() {
	espresso := starbuz.NewEspresso()
	fmt.Printf("%s: %.2f\n", espresso.GetDescription(), espresso.Cost())
	mocha := starbuz.NewMohca(espresso)
	fmt.Printf("%s: %.2f\n", mocha.GetDescription(), mocha.Cost())
	whipMocha := starbuz.NewWipe(mocha)
	fmt.Printf("%s: %.2f\n", whipMocha.GetDescription(), whipMocha.Cost())
}