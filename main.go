package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/time/:param", func(c *gin.Context) {
		// 获取param
		param := c.Param("param")
		// 拿到当前时间并初始化
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		// 拼接
		result := fmt.Sprintf("%s - %s", param, currentTime)
		// 返回
		c.String(200, result)
	})
	// 启动服务
	r.Run(":8080")
}
