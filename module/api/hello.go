package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Test struct {
}

func NewTest() Test {
	return Test{}
}

func (t Test) Get(ctx *gin.Context) {
	log.Println("Hello 接收到GET请求..")
	ctx.JSON(http.StatusOK, "已接收到GET请求..")
	return
}
