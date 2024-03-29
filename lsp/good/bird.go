package good

import "fmt"

// だめな例(is-aの関係がなりたっていない)
// こじれた技巧的な派生
// 基本クラスの一部分の機能しか使っていない＋追加機能を付けた派生

type Bird struct {
	name string
}

func (b *Bird) fly() {
	fmt.Printf("%sは飛びます\n", b.name)
}

type Crow struct {
	Bird
}

func NewCrow() *Crow {
	c := &Crow{Bird: Bird{name: "カラス"}}
	return c
}
