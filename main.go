package main

import (
	"blog-go/router"
	"log"
)

//TODO 补充日志打印到控制台
func main() {

	// 1.创建路由
	r := router.Init()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080"	)
	err := r.Run(":8000")
	if err != nil {
		log.Fatalf("Start server: %+v", err)
	}
}
