package main

import (
	"Nexus/internal/pkg"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化gin引擎
	r := gin.Default()

	//配置路由
	r.GET("/ping", func(c *gin.Context) {
		pkg.OK(c, "操作成功", nil)
	})
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("服务启动失败{}", err.Error())
		return
	}
}
