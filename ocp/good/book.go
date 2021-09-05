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
	return fmt.Sprintf("これは%sです\n", b.kind.String())
}

// [追加仕様]
// もしprintAllOfRentalTypeで表示順番をドメイン知識として制御したい(つまりこの関数の中で制御したい)
// この仕様に対してOCPは準拠していない(printRentalTypeの抽象化の意味がなかった)
// 結論: あらゆる変更に対して完璧に閉じることは不可能。設計担当者はどういった種類の変更が頻繁に起こるのかを予測しそれに対してOCP準拠しておくこと
func printAllOfRentalType(books []*Book) []string {
	prints := make([]string, len(books))
	for i, b := range books {
		prints[i] = b.printRentalType()
	}
	return prints
}
