package router

import (
	"Termbin/api"

	"github.com/gin-gonic/gin"
)

// InitRouter 设置路由
func InitRouter(r *gin.Engine) {
	r.POST("/", api.NewClipboard())
	r.GET("/:id", api.GetClipboard())
	r.PUT("/:id", api.UpdateClipboard())
	r.DELETE("/:id", api.DeleteClipboard())

	v1 := r.Group("api/v1")
	{
		v1.POST("user/sign-in", api.UserSignIn())
		v1.POST("user/sign-up", api.UserSignUp())
	}

}
