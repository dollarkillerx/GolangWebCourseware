package router

import (
	"GolangWebCourseware/dbops"
	"GolangWebCourseware/defs"
	"GolangWebCourseware/response"
	"GolangWebCourseware/utils"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)


// url GET /user/addUser
// 展示添加用户页面
func ShowAddUser(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	// 加载模板
	files, e := template.ParseFiles("./template/login.html")
	if e != nil {
		fmt.Println("html 解析出问题了")
		response.SendErrorResponse(w,defs.ErrorTemplateError)
		return
	}
	r.ParseForm()
	score := r.FormValue("score")
	i, _ := strconv.Atoi(score)
	// 往Template填充内容
	type Data struct {
		Title string
		Msg string
		Score int
		Maps map[int]string
	}
	maps := make(map[int]string)
	maps[0] = "hello"
	maps[1] = "golang"
	maps[2] = "你好啊"
	//e = files.Execute(w,&Data{
	//	Title:"注册",
	//	Msg:"html/Template学习",
	//	Score:i,
	//	Maps:maps,
	//})

	e = files.Execute(w,map[string]interface{}{
		"Request":r,
		"Score":i,
		"Msg":"html/Template学习",
		"Maps":maps,
	})
	if e != nil {
		fmt.Println("填充数据出问题了",e.Error())
		response.SendErrorResponse(w,defs.ErrorTemplateError)
		return
	}
}

// url POST /user/addUser
// 添加用户
func AddUser(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	r.ParseForm()
	name := r.Form["name"]
	password := r.Form["password"]
	if len(name) == 0 || len(password) == 0{
		response.SendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	user := &defs.User{Name:name[0],Password:password[0]}

	// 用户注册
	err := dbops.RegisterUser(user)
	if err != nil {
		response.SendErrorResponse(w,defs.ErrorDBError)
		return
	}

	response.SendNormalResponse(w,"success",http.StatusCreated)
}

func uploadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	//限制文件大小
	r.Body = http.MaxBytesReader(w,r.Body,defs.MAX_UPLOAD_SIZE)

	//这个是解析,顺便设置表单最大大小
	if err := r.ParseMultipartForm(defs.MAX_UPLOAD_SIZE);err != nil{
		response.SendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}

	file,_, e := r.FormFile("file")
	if e != nil {
		response.SendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}

	bytes, _ := ioutil.ReadAll(file) //读文件

	s, _ := utils.NewUUIDSimplicity()
	e = ioutil.WriteFile((defs.FILE_DIR + s), bytes, 0666)
	if e != nil {
		response.SendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}

}

func downloadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	vid :=p.ByName("vid-id")
	vl := defs.FILE_DIR + vid

	file, e := os.Open(vl)
	if e != nil{
		response.SendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}

	//w.Header().Set("Content-Type","")
	http.ServeContent(w,r,"",time.Now(),file)

	defer file.Close()
}