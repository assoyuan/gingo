package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIController struct{}

func (con APIController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"msg":    "API",
	})
}
