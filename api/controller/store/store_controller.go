package storecontroller

import "github.com/gin-gonic/gin"

type StoreController interface {
	Register(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Edit(ctx *gin.Context)
	Unreg(ctx *gin.Context)
}
