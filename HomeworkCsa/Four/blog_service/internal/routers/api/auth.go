package api

import (
	"blog_service.com/m/global"
	"blog_service.com/m/internal/service"
	"blog_service.com/m/pkg/app"
	"blog_service.com/m/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context)  {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Fatalf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Fatalf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.Appkey, param.AppSecret)
	if err != nil {
		global.Logger.Fatalf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
