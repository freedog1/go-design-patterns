package starbuz

// Beverage インタフェースの定義
type Beverage interface {
    GetDescription() string
    Cost() float64
}

// BeverageImpl 構造体の定義
type BeverageImpl struct {
    Description string
}

// GetDescription メソッドの実装
func (b *BeverageImpl) GetDescription() string {
    return b.Description
}

// Cost メソッドの実装
func (b *BeverageImpl) Cost() float64 {
    // サブクラスでの具体的な実装が必要
    return 0.0
}

// Espresso 具体的な飲み物
type Espresso struct{
	Description string
}

func NewEspresso() *Espresso {
	return &Espresso{
		Description: "エスプレッソ",
	}
}

func (e *Espresso) Cost() float64 {
	return 1.99
}

func (e *Espresso) GetDescription() string {
	return e.Description
}

type Milk struct{
	Description string
}

func NewMilk() *Milk {
	return &Milk{
		Description: "ミルク",
	}
}

func (m *Milk) Cost() float64 {
	return 1.00
}

func (m *Milk) GetDescription() string {
	return m.Description
}
