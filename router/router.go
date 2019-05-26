package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/**
创建路由注册
*/
func RegisterRouter() *httprouter.Router {
	router := httprouter.New() // 得到router实例

	router.GET("/user/addUser",ShowAddUser) // 展示添加用户页面
	router.POST("/user/addUser",AddUser) // 添加用户

	//静态资源
	router.ServeFiles("/static/*filepath",http.Dir("./template/static/"))

	return router
}
