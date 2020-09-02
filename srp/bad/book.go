package bad

import (
	"fmt"
	"time"
)

// Book Domain
type Book struct {
	title       string
	author      string
	publishDate time.Time
}

func (b *Book) Title() string {
	return b.title
}
func (b *Book) Author() string {
	return b.author
}
func (b *Book) PublishDate() time.TIme {
	return b.publishDate
}

// これだと保存処理をする場合、本の概要を変更する場合など複数の理由で変更が必要になってしまう可能性がある
func Register(b *Book) error {
	// saved book
	return nil
}
