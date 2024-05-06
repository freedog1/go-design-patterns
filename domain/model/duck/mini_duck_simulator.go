package duck

import "fmt"

type MallerdDuck struct {
	Duck *Duck
}

func MiniDuckSimulator() *MallerdDuck{
	return &MallerdDuck{
		Duck: &Duck{quickBehavior: &Quack{}, flyBehavior: &FlyWithWings{}},
	}
}

func (m *MallerdDuck) Display(){
	fmt.Println("マガモです")
}

func ModelDuckSimulator() *ModelDuck{
	return &ModelDuck{
		Duck: &Duck{quickBehavior: &Quack{}, flyBehavior: &FlyWithWings{}},
	}
}