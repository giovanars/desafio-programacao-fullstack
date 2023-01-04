package dto

import (
	"time"

	"github.com/giovanars/desafio-programacao-fullstack/entity"
)

type TransactionDto struct {
	ID          entity.ID `json:"id"`
	Type        string    `json:"type"`
	Date        time.Time `json:"date"`
	ProductName string    `json:"productName"`
	Value       float64   `json:"value"`
	SellerName  string    `json:"sellerName"`
}
