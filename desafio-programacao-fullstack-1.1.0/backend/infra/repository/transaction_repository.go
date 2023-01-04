package repository

import "github.com/giovanars/desafio-programacao-fullstack/entity"

type TransactionRepository struct {
}

func NewTransactionRepository() TransactionRepository {
	return TransactionRepository{}
}

func (repo TransactionRepository) Get(id entity.ID) (*entity.Transaction, error) {
	return nil, nil
}

func (repo TransactionRepository) List() ([]*entity.Transaction, error) {
	return nil, nil
}

func (repo TransactionRepository) Create(t *entity.Transaction) (entity.ID, error) {
	return entity.NewID(), nil
}

func (repo TransactionRepository) Update(t *entity.Transaction) error {
	return nil
}
