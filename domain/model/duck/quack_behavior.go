package duck

import "fmt"

type QuackBehavior interface {
	quack()
}

type Quack struct {
	QuackBehavior *QuackBehavior
}

func (q *Quack) quack() {
	fmt.Println("ガーガー")
}

type MuteQuack struct {
	QuackBehavior *QuackBehavior
}

func (q *MuteQuack) quack() {
	fmt.Println("キューキュー")
}
