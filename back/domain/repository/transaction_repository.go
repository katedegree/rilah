// REVIEW: トランザクションをリポジトリで管理するかどうか
package repository

type TransactionRepository interface {
	ExecuteWith(fn func() error) error
}
