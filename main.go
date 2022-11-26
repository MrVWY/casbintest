package main

import (
	"casbintest/global"
	"casbintest/module/model"
	"casbintest/module/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	gin.SetMode("debug")
	router := router.NewRouter()

	global.Enforcer = model.Casbin()
	//log.Fatal("开始启动")
	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("s.ListenAndServer err: %v", err)
	}
}
