package middleawar

import (
	"casbintest/module/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CasbinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//response := app.NewResponse(ctx)
		// 获取请求的URI
		obj := ctx.Request.URL.RequestURI()
		// 获取请求方法
		act := ctx.Request.Method

		// 获取当前用户的角色
		sub := "p1"
		// 获取当前用户所操作的的项目
		dom := "A123456"

		e := model.Casbin()
		fmt.Println(obj, dom, act, sub)

		// 判断策略中是否存在
		success := e.Enforce(sub, dom, obj, act)
		fmt.Println("判断策略", success)
		if success {
			log.Println("恭喜您,权限验证通过")
			ctx.Next()
		} else {
			log.Printf("e.Enforce err: %s", "很遗憾,权限验证没有通过")
			//response.ToErrorResponse(400, "很遗憾,权限验证没有通过")
			//ctx.Abort()
			ctx.JSON(http.StatusOK, "很遗憾,权限验证没有通过")
			ctx.Abort()
			//ctx.Next()
			return
		}
	}
}
