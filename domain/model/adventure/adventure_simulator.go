package adventure

func NewKing() *King {
	king := &King{
		&Character{
			WeaponBehavior: &SwordBehavior{},
		},
	}

	return king
}