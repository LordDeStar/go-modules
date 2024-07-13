package go-modules

type IDestructible interface {
	ShowHp()
	TakeDamage(damage uint8)
}
