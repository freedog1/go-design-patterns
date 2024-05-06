package adventure

import "fmt"

type WeaponBehavior interface {
	useWeapon()
}

type KnifeBehavior struct {
	WeapontBehavior *WeaponBehavior
}

func (b *KnifeBehavior) useWeapon() {
	fmt.Println("ナイフで切る")
}

type BowAndAllowBehavior struct {
	WeapontBehavior *WeaponBehavior
}

func (b *BowAndAllowBehavior) useWeapon() {
	fmt.Println("矢を引く")
}

type AxeBehavior struct {
	WeapontBehavior *WeaponBehavior
}

func (b *AxeBehavior) useWeapon() {
	fmt.Println("斧で叩き切る")
}

type SwordBehavior struct {
	WeapontBehavior *WeaponBehavior
}

func (b *SwordBehavior) useWeapon() {
	fmt.Println("剣を振り下ろす")
}
