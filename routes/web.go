package routes

import (
	"gingo/controllers"
	"github.com/gin-gonic/gin"
)

func InitWebRouter(e *gin.Engine) {
	webGroup := e.Group("/web")
	{
		webGroup.GET("/index", controllers.WebController{}.Index)
	}
}
