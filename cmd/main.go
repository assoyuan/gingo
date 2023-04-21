package main

import (
	"gingo/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//s := gin.New()

	s := gin.Default()

	//引入HTML模板
	s.LoadHTMLGlob("../templates/*")

	//注册路由
	routes.InitAPIRouter(s)
	routes.InitWebRouter(s)

	s.GET("/", func(ctx *gin.Context) {
		routers := make([]string, 0)

		for _, v := range s.Routes() {
			routers = append(routers, v.Method+" "+v.Path)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"routes": routers,
		})
	})

	_ = s.Run(":8088")
}
