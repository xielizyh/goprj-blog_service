package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/xielizyh/goprj-blog_service/global"
	"github.com/xielizyh/goprj-blog_service/internal/model"
	"github.com/xielizyh/goprj-blog_service/internal/routers"
	"github.com/xielizyh/goprj-blog_service/pkg/logger"
	"github.com/xielizyh/goprj-blog_service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	if err := setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	if err := setupLogger(); err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/xielizyh/goprj-blog_service
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
	// s.ListenAndServe()
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe error: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	// 接收syscall.SIGINT 和syscall.SIGTERM信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 最大时间控制，通知该服务端有5s的时间来处理原有的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		// 一旦5s超时，强制关闭
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
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
	// JWT配置
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	global.JWTSetting.Expire *= time.Second

	// 打印设置
	fmt.Println("settings:", global.ServerSetting, global.AppSetting, global.DatabaseSetting, global.JWTSetting)
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

func setupLogger() error {
	//  使用lumberjack 作为日志库的 io.Writer
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		// 日志文件名
		Filename: global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		// 日志最大占用空间600MB
		MaxSize: 600,
		// 日志文件最大生存周期为 10 天
		MaxAge: 10,
		// 日志文件名的时间格式为本地时间
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	// 测试日志存储
	global.Logger.Infof("%s: blog-service/%s", "xieli", "test logger")
	return nil
}
