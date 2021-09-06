package bad

type AnimalAction interface {
	swim()
	eat()
	fly()
}

// カラスはswimは必要ない
type CrowAction interface {
	AnimalAction
}

// クジラはflyは必要ない
type WhaleAction interface {
	AnimalAction
}
