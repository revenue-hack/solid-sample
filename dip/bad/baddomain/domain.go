package baddomain

import (
	"time"

	"github.com/revenue-hack/solid-sample/dip/bad"
)

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

func (b *Book) Register() error {
	// ドメインが技術的詳細に依存している
	return (&bad.BookRepositoryImpl{}).Register()
}
