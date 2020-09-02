package bad

import (
	"fmt"
	"time"
)

// Book Domain
type Book struct {
	title       string
	author      string
	kind        string
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

// 色々な本が増えていくたびにこの関数のロジックが追加される(修正に対して閉じていない)
func (b *Book) printRentalType() string {
	switch b.kind {
	case "novel":
		return "これは小説です"
	case "comics":
		return "これは漫画です"
	default:
		return ""
	}
}
