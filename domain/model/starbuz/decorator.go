package starbuz

// CondimentDecorator インタフェースの定義
type CondimentDecorator interface {
	GetDescription() string
}

type Mocha struct{
	Beverage Beverage
}

func NewMohca(beverage Beverage) *Mocha {
	return &Mocha{
		Beverage: beverage,
	}
}

func (m *Mocha) GetDescription() string {
	return m.Beverage.GetDescription() + ". モカ"
}

func (m *Mocha) Cost() float64 {
	return m.Beverage.Cost() + 0.2
}

type Whipe struct {
	Beverage Beverage
}

func NewWipe(beverage Beverage) *Whipe {
	return &Whipe{
		Beverage: beverage,
	}
}

func (m *Whipe) GetDescription() string {
	return m.Beverage.GetDescription() + ". ホイップ"
}

func (m *Whipe) Cost() float64 {
	return m.Beverage.Cost() + 0.1
}
