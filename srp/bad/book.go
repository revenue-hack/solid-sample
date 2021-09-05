package bad

import (
	"time"
)

// 悪い理由: ビジネスルールと永続性のあるシステムを包含してしまっている
// ビジネスルールは頻繁に変更されるものだが、永続性のあるシステムはされない
// また永続性のあるシステムはビジネスルールとは違った理由で変更する
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
	// 永続化処理
	// saved book
	return nil
}
