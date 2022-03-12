package good

type SwimAction interface {
	swim()
}
type EatAction interface {
	eat()
}
type FlyAction interface {
	fly()
}

// 実装する必要のないものは実装しない

type CrowAction interface {
	EatAction
	FlyAction
}

type WhaleAction interface {
	EatAction
	SwimAction
}
