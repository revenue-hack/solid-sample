package bad

import (
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
func (b *Book) PublishDate() time.Time {
	return b.publishDate
}

func (b *Book) printRentalType() string {
	switch b.kind {
	case "novel":
		return "これは小説です"
	case "comics":
		return "これは漫画です"
	case "documentary":
		return "これはドキュメンタリです"
	default:
		return ""
	}
}
