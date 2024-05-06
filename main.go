package main

import "design-pattern/domain/model/duck"

func main() {
	miniDuck := duck.MiniDuckSimulator()
	miniDuck.Display()
	miniDuck.Duck.PerformQuack()
	miniDuck.Duck.PerformFly()

	modelDuck := duck.ModelDuckSimulator()
	modelDuck.Display()
	modelDuck.Duck.SetFlyBehavior(&duck.FlyRocketPowered{})
	modelDuck.Duck.PerformFly()
}
