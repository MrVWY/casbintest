package router

import (
	"casbintest/middleawar"
	v12 "casbintest/module/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	casbin := v12.NewCasbin()
	test := v12.NewTest()
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleawar.CasbinHandler())

	{
		// 测试路由
		apiv1.GET("/hello", test.Get)

		// 权限策略管理
		apiv1.POST("/casbin", casbin.Create)
		apiv1.POST("/casbin/list", casbin.List)
	}
	return r
}
