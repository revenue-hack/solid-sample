package good

type SwimAction interface {
	swim()
}
type EastAction interface {
	eat()
}
type FlyAction interface {
	fly()
}

// 実装する必要のないものは実装しない

type CrowAction interface {
	EastAction
	FlyAction
}

type WhaleAction interface {
	EastAction
	SwimAction
}
