package v1

import (
	"blog_service.com/m/pkg/app"
	"blog_service.com/m/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{

 }
func NewArticle() Article {
 	return Article{}
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

// List
//@Summary 获取单个文章
//@Produce json
//@Param name body string true "文章名称" minlength(3) maxlength(100)
//@Param state body int false "状态" Enums(0, 1) default(1)
//@Param created_by body string true "创建者" minlength(3) maxlength(100)
//@Success 200 {object} model.Tag "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/articles [post]
func (a Article) List(c *gin.Context) {}

//Create
//@Summary 创建文章
//@Produce json
//@Param name body string true "文章名称" minlength(3) maxlength(100)
//@Param state body int false "状态" Enums(0, 1) default(1)
//@Param created_by body string true "创建者" minlength(3) maxlength(100)
//@Success 200 {object} model.Tag "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {}

// Update
// @Summary 更新文章
// @Produce json
// @Param id path int true "文章 ID"
// @Param name body string false "文章名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {}

// Delete
// @Summary 删除文章
// @Produce json
// @Param id path int true "文章 ID"
// @Success 200 {object} model.TagSwagger "成功
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {}
