package good

// 保存処理を分離することで本のドメインルールの変更に影響しなくなる
type BookRepository interface {
	Register() (*Book, error)
}
