package structs

type Player struct {
	Name           string
	Point          int
	Dice           []int
	DiceFromPlayer []int
	IsActive       bool
}