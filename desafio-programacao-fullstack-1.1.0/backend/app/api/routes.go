package api

import (
	"github.com/gin-gonic/gin"
	"github.com/giovanars/desafio-programacao-fullstack/app/api/controllers"
	"github.com/giovanars/desafio-programacao-fullstack/app/api/middlewares"
	"github.com/giovanars/desafio-programacao-fullstack/config"
	"github.com/giovanars/desafio-programacao-fullstack/infra/repository"
)

type Router struct {
}

func Routers() Endpoint {
	return &Router{}
}

func (handler *Router) Handler(router gin.IRouter, config *config.Config, server *Server) error {

	router.GET("/health", middlewares.GinHealthHandler("/health"))

	transactionController := handler.transactionController(server)
	router.GET("v1/transactions", transactionController.GetAll)

	uploadController := handler.uploadController(server)
	router.POST("v1/upload", uploadController.Upload)

	return nil
}

func (handler *Router) transactionController(server *Server) *controllers.TransactionController {
	return controllers.NewTransactionController(
		repository.NewTransactionRepository(),
	)
}

func (handler *Router) uploadController(server *Server) *controllers.UploadController {
	return controllers.NewUploadController()
}
