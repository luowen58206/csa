package routers

import (
	"blog_service.com/m/global"
	"blog_service.com/m/internal/service"
	"blog_service.com/m/pkg/app"
	"blog_service.com/m/pkg/convert"
	"blog_service.com/m/pkg/errcode"
	"blog_service.com/m/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct {

}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context)  {
		response := app.NewResponse(c)
		file, fileHeader, err := c.Request.FormFile("file")
		if err != nil {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
			return
		}

		fileType := convert.StrTo(c.PostForm("type")).MustInt()
		if fileHeader == nil || fileType <= 0 {
			response.ToErrorResponse(errcode.InvalidParams)
			return
		}

		svc := service.New(c.Request.Context())
		fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
		if err != nil {
			global.Logger.Fatalf("svc.UploadFile err: %v",err)

			response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
			return
		}

		response.ToResponse(gin.H{
			"file_access_url": fileInfo.AccessUrl,
		})
}
