package main

import (
	"net/http"
	"time"

	"github.com/xielizyh/goprj-blog_service/internal/routers"
)

func main() {
	// 新建路由
	router := routers.NewRouter()
	// 定义http.Server
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 开始监听
	s.ListenAndServe()
}
