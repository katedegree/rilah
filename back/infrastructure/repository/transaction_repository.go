package repository

import (
	"back/domain/repository"

	"gorm.io/gorm"
)

type transactionRepository struct {
	orm *gorm.DB
}

func NewTransactionRepository(orm *gorm.DB) repository.TransactionRepository {
	return &transactionRepository{
		orm: orm,
	}
}

func (r *transactionRepository) ExecuteWith(fn func() error) error {
	tx := r.orm.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
