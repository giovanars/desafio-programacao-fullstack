package api

import "github.com/giovanars/desafio-programacao-fullstack/usecase/transaction"

type Service struct {
	TransactionService transaction.TransactionService
}

func NewService(
	transactionService transaction.TransactionService,
) *Service {
	return &Service{
		TransactionService: transactionService,
	}
}
