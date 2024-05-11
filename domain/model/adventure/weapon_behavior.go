package adventure

import "fmt"

type WeaponBehavior interface {
	useWeapon()
}

type KnifeBehavior struct {
}

func (b *KnifeBehavior) useWeapon() {
	fmt.Println("ナイフで切る")
}

type BowAndAllowBehavior struct {
}

func (b *BowAndAllowBehavior) useWeapon() {
	fmt.Println("矢を引く")
}

type AxeBehavior struct {
}

func (b *AxeBehavior) useWeapon() {
	fmt.Println("斧で叩き切る")
}

type SwordBehavior struct {
}

func (b *SwordBehavior) useWeapon() {
	fmt.Println("剣を振り下ろす")
}
