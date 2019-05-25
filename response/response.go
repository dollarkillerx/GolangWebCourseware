package response

import (
	"GolangWebCourseware/defs"
	"encoding/json"
	"net/http"
)

// 当返回错误信息时
func SendErrorResponse(w http.ResponseWriter,response defs.ErrResponse)  {
	w.WriteHeader(response.HttpSc)
	bytes, _ := json.Marshal(response.Error)
	w.Write(bytes)
}

// 返回自定义类型
func SendNormalResponse(w http.ResponseWriter,resp string,sc int) {
	w.WriteHeader(sc)
	w.Write([]byte(resp))
}