package duck

import "fmt"

type FlyBehavior interface {
	fly()
}

type FlyWithWings struct {
	FlyBehavior *FlyBehavior
}

func (f *FlyWithWings) fly() {
	fmt.Println("飛んでいる")
}

type FlyNoWay struct {
	FlyBehavior *FlyBehavior
}

func (f *FlyNoWay) fly() {
	fmt.Println("飛べない")
}

type FlyRocketPowered struct {
	FlyBehavior *FlyBehavior
}

func (f *FlyRocketPowered) fly() {
	fmt.Println("ロケットで飛ぶ")
}