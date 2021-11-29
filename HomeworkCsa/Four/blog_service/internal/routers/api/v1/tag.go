package v1

import (
	"blog_service.com/m/global"
	"blog_service.com/m/internal/service"
	"blog_service.com/m/pkg/app"
	"blog_service.com/m/pkg/convert"
	"blog_service.com/m/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct {

}

func NewTag() Tag {
	return Tag{}
}

// Get 获取tag
func (t Tag) Get(c *gin.Context) {}

// List @Summary 获取多个标签
//@Produce json
//@Param name query string false "标签名称" maxlength(100)
//@Param state query int false "状态" Enums(0, 1) default(1)
//@Param page query int false "页码"
//@Param page_size query int false "每页数量"
//@Success 200 {object} model.Tag "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Fatalf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

		/*
		我们完成了获取标签列表接口的处理方法
		我们在方法中完成了入参校验和绑定、
		获取 标签总数、获取标签列表、 序列化结果集
		四大功能板块的逻辑串联和日志、错误处理。
		 */
		svc := service.New(c.Request.Context())
		pager := app.Pager{
			Page: app.GetPage(c),
			PageSize: app.GetPageSize(c),
		}
		totalRows, err := svc.CountTag(&service.CountTagRequest{
			Name: param.Name,
			State: param.State,
		})
		if err != nil {
			global.Logger.Fatalf("svc.CountTag err: %v",err)
			response.ToErrorResponse(errcode.ErrorCountTagFail)
			return
		}

		tags, err := svc.GetTagList(&param, &pager)
		if err != nil {
			global.Logger.Fatalf("svc.GetTagList err: %v",err)
			return
		}

		response.ToResponseList(tags, totalRows)
		return
}

// Create
// @Summary 新增标签
//@Produce json
//@Param name body string true "标签名称" minlength(3) maxlength(100)
//@Param state body int false "状态" Enums(0, 1) default(1)
//@Param created_by body string true "创建者" minlength(3) maxlength(100)
//@Success 200 {object} model.Tag "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Fatalf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Fatalf("svc.CreateTag err: %v",err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	response.ToResponse(gin.H{
		"message": "add success",
	})
	return
}

// Update
// @Summary 更新标签
// @Produce json
// @Param id path int true "标签 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Fatalf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Fatalf("svc.UpdateTag err: %v",err)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// Delete
// @Summary 删除标签
// @Produce json
// @Param id path int true "标签 ID"
// @Success 200 {object} model.TagSwagger "成功
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {

	param := service.DeleteTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Fatalf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Fatalf("svc.DeleteTag err: %v",err)
		return
	}

	response.ToResponse(gin.H{})
	return
}