package gomodules

type IDestructible interface {
	ShowHp()
	TakeDamage(damage uint8)
}
