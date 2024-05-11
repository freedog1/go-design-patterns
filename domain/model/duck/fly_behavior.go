package duck

import "fmt"

type FlyBehavior interface {
	fly()
}

type FlyWithWings struct {
}

func (f *FlyWithWings) fly() {
	fmt.Println("飛んでいる")
}

type FlyNoWay struct {
}

func (f *FlyNoWay) fly() {
	fmt.Println("飛べない")
}

type FlyRocketPowered struct {
}

func (f *FlyRocketPowered) fly() {
	fmt.Println("ロケットで飛ぶ")
}