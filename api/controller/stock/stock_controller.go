package stockcontroller

import "github.com/gin-gonic/gin"

type StockController interface {
	Register(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Edit(ctx *gin.Context)
	Unreg(ctx *gin.Context)
}
