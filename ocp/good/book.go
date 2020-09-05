package good

import (
	"fmt"
	"time"
)

// Book Domain
type Book struct {
	title       string
	author      string
	kind        BookKind
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

// 色々な本が増えていくたびにこの関数のロジックが追加される(修正に対して閉じていない)
// 本の種類が増えても修正する必要がなくなった
func (b *Book) printRentalType() string {
	return fmt.Sprintf("これは%sです", b.kind.String())
}
