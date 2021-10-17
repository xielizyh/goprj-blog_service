package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/xielizyh/goprj-blog_service/global"
	"github.com/xielizyh/goprj-blog_service/internal/model"
	"github.com/xielizyh/goprj-blog_service/internal/routers"
	"github.com/xielizyh/goprj-blog_service/pkg/setting"
)

func init() {
	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	if err := setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func main() {
	// 新建路由
	router := routers.NewRouter()
	// 定义http.Server
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	// 开始监听
	s.ListenAndServe()
}

func setupSetting() error {
	// 新建设置
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	// 服务器配置
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	// App配置
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	// 数据库配置
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	fmt.Println("settings:", global.ServerSetting, global.AppSetting, global.DatabaseSetting)
	return nil
}

func setupDBEngine() error {
	var err error
	// 注意这里不是 “:=”，否则global.DBEngine得不到值
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
