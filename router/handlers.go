package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Home(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	w.Write([]byte("this is home page"))
}

// url: http:127.0.0.1:8085/name/:name
// p httprouter.Params 可以通过 p.ByName("name") 来获取值
// 现在我们来编译测试一下吧
func Name(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	name := p.ByName("name")

	w.Write([]byte("you name:"+name))
}