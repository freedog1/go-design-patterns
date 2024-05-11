package duck

import "fmt"

type QuackBehavior interface {
	quack()
}

type Quack struct {
}

func (q *Quack) quack() {
	fmt.Println("ガーガー")
}

type MuteQuack struct {
}

func (q *MuteQuack) quack() {
	fmt.Println("キューキュー")
}
