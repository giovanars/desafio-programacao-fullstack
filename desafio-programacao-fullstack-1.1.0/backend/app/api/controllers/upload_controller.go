package controllers

import "github.com/gin-gonic/gin"

type UploadController struct {
}

func NewUploadController() *UploadController {
	return &UploadController{}
}

func (controller *UploadController) Upload(context *gin.Context) {

}
