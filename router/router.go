package router

import (
	"github.com/julienschmidt/httprouter"
)

/**
创建路由注册
*/
func RegisterRouter() *httprouter.Router {
	router := httprouter.New() //得到router实例

	router.GET("/name/:name",Name)

	router.GET("/registeredusers",RegisteredUsersGet)
	router.POST("/registeredusers",RegisteredUsersPOST)
	router.POST("/json",RegisteredUsersPOSTByJSON)

	return router
}
