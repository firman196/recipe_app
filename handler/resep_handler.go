package handler

import "github.com/gin-gonic/gin"

type ResepHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetById(c *gin.Context)
	Delete(c *gin.Context)
}
