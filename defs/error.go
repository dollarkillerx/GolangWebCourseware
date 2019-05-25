package defs

import "net/http"

type Err struct {
	Error string `json:"error"`  //这是打tag
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSc int
	Error Err
}

// 定义常用返回信息
var (
	// 定义参数错误返回类型
	ErrorRequestBodyParseFailed = ErrResponse{Error:Err{Error:"Request body is not correct",ErrorCode:"001"},HttpSc:http.StatusBadRequest}
)