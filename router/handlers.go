package router

import (
	"GolangWebCourseware/defs"
	"GolangWebCourseware/response"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
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

// url: get http:127.0.0.1:8085/registeredusers
// 演示get请求获取数据
func RegisteredUsersGet(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	r.ParseForm() // 解析url传递的参数，对于POST则解析响应包的主体（request body）
	// 注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	name := r.Form["name"]
	password := r.Form["password"]
	// 这里获取到的是[]string

	if len(name)==0 || len(password)==0 {
		// 如果没有传入 就访问这个url 处理
		// 这里我先不封装,后面的课程会讲分离
		// RESTful知道吧  参数异常返回400
		response.SendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}

	w.Write([]byte("name: "+name[0]+" password: "+password[0]))
	// 此处小朋友会疑惑吗name[0]  接收的参数为什么是[]string

	/**
	我来解释

	用postman测试，提交http://localhost:8080/?uid=111
	服务端输出 ：[111]
	提交： http://localhost:8080/?uid=111&uid=222
	服务端输出：[111 222]
	 */
}

// url: post http:127.0.0.1:8085/registeredusers
// 演示post请求获取数据
func RegisteredUsersPOST(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	/**
	注意.Form
	r.Form包含了GET、POST参数
	POST:为 application/x-www-form-urlencoded 提交的蛤
	 */

	r.ParseForm()
	name := r.Form["name"]
	password := r.Form["password"]

	if len(name) == 0 || len(password) == 0 {
		response.SendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	w.Write([]byte("name: "+name[0]+" password: "+password[0]))
}

// url: post http:127.0.0.1:8085/rpjson
func RegisteredUsersPOSTByJSON(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	//读取body 这里就用ioutil包吧,这个没有包含文件不会太大
	bytes, _ := ioutil.ReadAll(r.Body)
	// 创建一个用户接收数据的空对象
	userinfo := &defs.User{}

	//序列化
	err := json.Unmarshal(bytes, userinfo)

	if err != nil {
		response.SendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}

	fmt.Println(userinfo)
	w.Write(bytes)
}