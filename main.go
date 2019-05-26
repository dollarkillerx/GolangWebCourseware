package main

import (
	"GolangWebCourseware/router"
	"fmt"
	"net/http"
)


func main() {
	registerRouter := router.RegisterRouter()// 注册路由

	fmt.Println("server is runing ...")

	// middleware 接替
	handler := router.NewMiddleWareHandler(registerRouter,1000)

	//这里发现改变了什么吗?
	//对第二次参数变成了router
	err := http.ListenAndServe(":8085", handler) //第一个是地址(ip:端口 ip可以省略 监听本机全部端口) 第二个是handler

	if err != nil {
		fmt.Println("server error:",err.Error())
	}
}
