package bad

import (
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
func (b *Book) PublishDate() time.Time {
	return b.publishDate
}

func Register(b *Book) error {
	// saved book
	return nil
}
