package handler

import "github.com/gin-gonic/gin"

type BahanHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetById(c *gin.Context)
}
