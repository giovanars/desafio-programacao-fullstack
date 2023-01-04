package transaction

import (
	"time"

	"github.com/giovanars/desafio-programacao-fullstack/entity"
)

type UseCase interface {
	GetTransaction(id entity.ID) (*entity.Transaction, error)
	GetAll() ([]*entity.Transaction, error)
	CreateTransaction(typeId entity.ID, date time.Time, productName string, Value float64, sellerName string) (entity.ID, error)
}

type Repositry interface {
	Get(id entity.ID) (*entity.Transaction, error)
	List() ([]*entity.Transaction, error)
	Create(t *entity.Transaction) (entity.ID, error)
	Update(t *entity.Transaction) error
}
