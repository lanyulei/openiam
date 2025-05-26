package router

import (
	"net/http"
	v1 "openops/common/router/v1"
	"time"

	"github.com/lanyulei/toolkit/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// Setup 加载路由
func Setup(g *gin.Engine) {
	// 使用zap接收gin框架默认的日志并配置
	g.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 静态文件打包进二进制文件中
	//g.GET("/assets/*filepath", func(c *gin.Context) {
	//	staticServer := http.FileServer(http.FS(assets.StaticFs))
	//	staticServer.ServeHTTP(c.Writer, c.Request)
	//})
	//g.SetHTMLTemplate(template.Must(template.New("").ParseFS(assets.TemplateFs, "template/*")))
	//g.Handle("GET", "/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", gin.H{})
	//})

	// 静态文件配置
	g.StaticFS("/static/uploadfile", http.Dir("static/uploadfile"))
	//g.StaticFile("/favicon.ico", "./static/assets/template/favicon.ico")

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "404 page not found",
		})
	})

	// pprof router
	pprof.Register(g)

	// cors， 跨域
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowOrigins:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	g.Use(cors.New(config))

	// 路由版本
	v1.RegisterRouter(g.Group(ApiV1Version))
}
