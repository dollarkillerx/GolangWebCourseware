package router

import (
	"GolangWebCourseware/response"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandle struct {
	R *httprouter.Router
	L *BucketLimit // 流控
}

func NewMiddleWareHandler(r *httprouter.Router,cc int) http.Handler {
	handle := &middleWareHandle{}
	handle.R =r
	handle.L = NewBucketLimit(cc)
	return handle
}

// 流控核心
func (m *middleWareHandle) ServeHTTP(w http.ResponseWriter,r *http.Request) () {
	// 从桶中获得令牌
	if !m.L.GetConn() {
		response.SendNormalResponse(w,"Too many requests",http.StatusTooManyRequests)
		return
	}
	m.R.ServeHTTP(w,r)
	defer func() {
		// 当连接结束 令牌返回令牌桶中
		m.L.ReleaseConn()
	}()
}
