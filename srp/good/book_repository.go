package good

// これだと保存処理をする場合、本の概要を変更する場合など複数の理由で変更が必要になってしまう可能性がある
// 保存処理を分離することで本の概要の変更に影響しなくなる
type BookRepository interface {
	Register() (*Book, error)
}
