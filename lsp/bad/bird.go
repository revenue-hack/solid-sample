package bad

import "fmt"

type Bird struct {
	name string
}

func (b *Bird) fly() {
	fmt.Printf("%sは飛びます\n", b.name)
}

type Whale struct {
	Bird
}

func NewWhale() *Whale {
	w := &Whale{Bird: Bird{name: "クジラ"}}
	return w
}

func (w *Whale) fly() {
	fmt.Printf("%sは飛びません\n", w.name)
}
