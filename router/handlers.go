package router

import (
	"GolangWebCourseware/dbops"
	"GolangWebCourseware/defs"
	"GolangWebCourseware/response"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func AddUser(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	bytes, _ := ioutil.ReadAll(r.Body)
	user := &defs.User{}

	err := json.Unmarshal(bytes, user)
	if err != nil {
		response.SendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}

	// 用户注册
	err = dbops.RegisterUser(user)
	if err != nil {
		response.SendErrorResponse(w,defs.ErrorDBError)
		return
	}

	response.SendNormalResponse(w,"success",http.StatusCreated)
}