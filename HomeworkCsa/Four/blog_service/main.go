package main

import (
	"blog_service.com/m/global"
	"blog_service.com/m/internal/model"
	"blog_service.com/m/internal/routers"
	"blog_service.com/m/pkg/logger"
	"blog_service.com/m/pkg/setting"
	"blog_service.com/m/pkg/tracer"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

/*
	·通过自定义 http.server、设置监听了 TCP Endpoint、处理的程序
	·允许读取/写入的最大时间、请求头的最大字节数等基础参数
	·最后调用了 ListenAndServer 方法开始监听
 */

func init()  {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err : %v",err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err : %v",err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err:%v",err)
	}

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v",err)
	}

}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr: ":" + global.ServerSetting.HttpPort,
		Handler: router,
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

	global.Logger.Infof("%s:blog_service.com/m/%s","eddycjy","blog-service")
}


func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server",&global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App",&global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database",&global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email",&global.EmailSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine,err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize: 600,
		MaxAge:  10,
		LocalTime: true,
	},"",log.LstdFlags).WithCaller(2)
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog_service",
		"127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}