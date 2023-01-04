package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/giovanars/desafio-programacao-fullstack/usecase/transaction"
)

type TransactionController struct {
	transactionRepository transaction.Repositry
}

func NewTransactionController(
	transactionRepository transaction.Repositry,
) *TransactionController {
	return &TransactionController{
		transactionRepository: transactionRepository,
	}
}

func (controller *TransactionController) GetAll(context *gin.Context) {

}
