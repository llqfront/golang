package router

import (
	. "blog.xhanglu.cn/create"
	. "blog.xhanglu.cn/get"
	"github.com/gin-gonic/gin"
)

func registerAPIRouter() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	routerDatasource := router.Group("/api/")
	routerDatasource.GET("/userinfo", Get)
	routerDatasource.POST("/add", Create)
	router.Run(":9502")

}
func Router() {
	registerAPIRouter()
}
