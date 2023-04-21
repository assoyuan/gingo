package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebController struct{}

func (con WebController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Web",
		"body":  "Web",
	})
}
