package bad

type Bird struct {
	name string
}

func (b *Bird) fly() {
	fmt.Printf("%sは飛びます\n", b.name)
}

type Whale struct {
	Bird
}

func NewWhale() *Crow {
	w := &Whale{name: "クジラ"}
	return w
}

func (w *Whale) fly() {
	fmt.Printf("%sは飛びません\n", w.name)
}
