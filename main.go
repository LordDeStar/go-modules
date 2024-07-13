package myModules

type IDestructible interface {
	ShowHp()
	TakeDamage(damage uint8)
}
