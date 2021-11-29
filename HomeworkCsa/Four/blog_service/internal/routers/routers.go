package routers

import (
	_ "blog_service.com/m/docs"
	"blog_service.com/m/global"
	"blog_service.com/m/internal/middleware"
	"blog_service.com/m/internal/routers/api"
	"blog_service.com/m/internal/routers/api/v1"
	"blog_service.com/m/pkg/limiter"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key: "/auth",
	FillInterval: time.Second,
	Capacity: 10,
	Quantum: 10,
})

func NewRouter() *gin.Engine  {

	//定义一个新的gin变量
	r := gin.New()

	//没有使用gin.Default()、所以需要使用内置函数logger（）和recovery（）
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.Recovery())
		r.Use(middleware.AccessLog())
	}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))
	r.Use(middleware.Translations())
	r.Use(middleware.Tracing())

	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/auth",api.GetAuth)
	//定义article和tag变量
	article := v1.NewArticle()
	tag := v1.NewTag()

	upload := NewUpload()
	r.POST("/upload/file",upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	//定义路由组、发起各种对文章的请求
	apiv1 :=r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		//新增标签
		apiv1.POST("/tags",tag.Create)
		//删除标签
		apiv1.DELETE("/tags/:id",tag.Delete)
		//更新指定标签的整个资源
		apiv1.PUT("/tags/:id",tag.Update)
		//更新指定标签的资源的某个字段
		apiv1.PATCH("/tags/:id/state",tag.Update)
		//
		apiv1.GET("/tags",tag.List)

		//新增文章
		apiv1.POST("/articles",article.Create)
		//删除文章
		apiv1.DELETE("articles/:id",article.Delete)
		//更新指定的文章
		apiv1.PUT("/articles/:id",article.Update)
		//更新文章中某个字段的内容
		apiv1.PATCH("/articles/:id/state",article.Update)
		//获取指定id的文章内容
		apiv1.GET("/articles/:id",article.Get)
		//获取整个文章列表的内容
		apiv1.GET("/articles",article.List)
	}

	return r


}
