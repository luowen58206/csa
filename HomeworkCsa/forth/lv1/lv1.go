package lv1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User string `json:"user" form:"user"`
	Password string `json:"password" form:"password"`
}

func Lv1()  {
	r := gin.Default()

	r.POST("/loginJson", func(c *gin.Context) {
		var login Login

		err := c.ShouldBind(&login)
		if err == nil {
			c.JSON(http.StatusOK,gin.H{
				"user": login.User,
				"password": login.Password,
			})
		}else {
			c.JSON(http.StatusBadRequest,gin.H{
				"error": err.Error(),
			})
		}
	})

	r.POST("/loginForm", func(c *gin.Context) {
		var login Login

		err := c.ShouldBind(&login)
		if err == nil {
			c.JSON(http.StatusOK,gin.H{
				"user": login.User,
				"password": login.Password,
			})
		}else {
			c.JSON(http.StatusBadRequest,gin.H{
				"error": err.Error(),
			})
		}
	})

	r.GET("/loginForm", func(c *gin.Context) {
		var login Login

		err := c.ShouldBind(&login)
		if err == nil {
			c.JSON(http.StatusOK,gin.H{
				"user": login.User,
				"password": login.Password,
			})
		}else {
			c.JSON(http.StatusBadRequest,gin.H{
				"error": err.Error(),
			})
		}
	})

	_ = r.Run(":9998")

}
