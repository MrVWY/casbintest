package v1

import (
	"casbintest/module/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Casbin struct {
}

func NewCasbin() Casbin {
	return Casbin{}
}

// Create godoc
// @Summary 新增权限
// @Description 新增权限
// @Tags 权限管理
// @Produce json
// @Security ApiKeyAuth
// @Param body service.CasbinCreateRequest true "body"
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/casbin [post]
func (c Casbin) Create(ctx *gin.Context) {
	param := service.CasbinCreateRequest{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	// 进行插入操作
	err = service.CasbinCreate(&param)
	if err != nil {
		log.Printf("svc.CasbinCreate err: %v", err)
		ctx.JSON(400, err)
	}
	ctx.JSON(200, "")
	return
}

// List godoc
// @Summary 获取权限列表
// @Produce json
// @Tags 权限管理
// @Security ApiKeyAuth
// @Param data body service.CasbinListRequest true "角色ID"
// @Success 200 {object} service.CasbinListResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/casbin/list [post]
func (c Casbin) List(ctx *gin.Context) {
	param := service.CasbinListRequest{}
	err := ctx.ShouldBind(&param)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	// 业务逻辑处理
	casbins := service.CasbinList(&param)
	var respList []service.CasbinInfo
	for _, host := range casbins {
		respList = append(respList, service.CasbinInfo{
			Path:   host[1],
			Method: host[2],
		})
	}
	ctx.JSON(http.StatusOK, respList)
	return
}
