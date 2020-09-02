package good

// 保存処理を分離することで本の概要の変更に影響しなくなる
type BookRepository interface {
	Register() (*Book, error)
}
