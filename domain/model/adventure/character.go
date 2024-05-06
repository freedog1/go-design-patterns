package adventure

type Character struct {
	WeaponBehavior WeaponBehavior
}

func (c *Character) UseWeapon() {
	c.WeaponBehavior.useWeapon()
}

type Queen struct {
	Character *Character
}

type King struct {
	Character *Character
}

type Knight struct {
	Character *Character
}

type Troll struct {
	Character *Character
}

func (c *Character) SetWeapon(w WeaponBehavior){
	c.WeaponBehavior = w	
}