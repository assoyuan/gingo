package main

import (
	"embed"
	"gingo/config"
	"gingo/routes"
	"gingo/utils"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//go:embed templates
var f embed.FS

func init() {
	// 注册配置文件
	config.InitConfig()

	// 注册日志组件
	utils.InitLog()
}

func main() {
	gin.SetMode(config.Config.App.Env)

	s := gin.Default()

	// 引入HTML模板
	temple := template.Must(template.New("").ParseFS(f, "templates/*"))
	s.SetHTMLTemplate(temple)

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
