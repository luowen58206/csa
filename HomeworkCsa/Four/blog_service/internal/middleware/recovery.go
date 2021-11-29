package middleware

import (
	"blog_service.com/m/global"
	"blog_service.com/m/pkg/app"
	"blog_service.com/m/pkg/email"
	"blog_service.com/m/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Recovery() gin.HandlerFunc {

	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host: global.EmailSetting.Host,
		Port: global.EmailSetting.Port,
		IsSSL: global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		Form: global.EmailSetting.Form,
	})

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil{
				global.Logger.WithCallersFrames().Fatalf("panic recover err: %v", err)

				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出,发送时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v",err),
					)
				if err != nil {
					global.Logger.Fatalf("mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
