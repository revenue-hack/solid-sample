package gooddomain

import "time"

// ドメイン
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
func (b *Book) PublishDate() time.Time {
	return b.publishDate
}

func (b *Book) Register(in BookDomainInterface) error {
	// ドメインのインターフェースに依存させることで依存方向を逆転している
	return in.Register()
}
