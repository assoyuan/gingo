package routes

import (
	"gingo/controllers"
	"github.com/gin-gonic/gin"
)

func InitAPIRouter(e *gin.Engine) {
	apiGroup := e.Group("/api")
	{
		apiGroup.GET("/index", controllers.APIController{}.Index)
	}
}
