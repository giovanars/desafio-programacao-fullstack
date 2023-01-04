package transaction

import (
	"time"

	"github.com/giovanars/desafio-programacao-fullstack/entity"
)

type TransactionService struct {
	abc string
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		abc: "",
	}
}

func (t TransactionService) GetTransaction(id entity.ID) (*entity.Transaction, error) {
	return nil, nil
}

func (t TransactionService) GetAll() ([]*entity.Transaction, error) {
	return nil, nil
}

func (t TransactionService) CreateTransaction(typeId entity.ID, date time.Time, productName string, Value float64, sellerName string) (entity.ID, error) {
	return entity.NewID(), nil
}
