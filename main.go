package main

import (
	"gingo/config"
	"gingo/routes"
	"gingo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	// 注册配置文件
	config.InitConfig()

	// 注册日志组件
	utils.InitLog()
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//s := gin.New()

	s := gin.Default()

	// 引入HTML模板
	s.LoadHTMLGlob("templates/*")

	// 注册路由
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

	_ = s.Run(config.Config.App.Addr)
}
