package duck

import "fmt"

type Duck struct {
	flyBehavior FlyBehavior
	quickBehavior QuackBehavior
}

func (d *Duck) Swim() {
	fmt.Println("すべてのカモは浮かぶ")
}

func (d *Duck) Display() {}

func (d *Duck) PerformQuack() {
	d.quickBehavior.quack()
}

func (d *Duck) PerformFly() {
	d.flyBehavior.fly()
}

func (d *Duck) SetFlyBehavior(flyBehavior FlyBehavior) {
	d.flyBehavior = flyBehavior
}

func (d *Duck) SetQuackBehavior(quickBehavior QuackBehavior) {
	d.quickBehavior = quickBehavior
}

type ModelDuck struct {
	Duck *Duck
}

func (d *ModelDuck) Display() {
	fmt.Println("模型のカモです")
}