package response

import (
	"GolangWebCourseware/defs"
	"encoding/json"
	"net/http"
)

// 当返回错误信息时
func SendErrorResponse(w http.ResponseWriter,errResp defs.ErrResponse)  {
	w.WriteHeader(errResp.HttpSc)
	bytes, _ := json.Marshal(errResp.Error)
	w.Write(bytes)
}