package router

import (
	"Termbin/api"
	"Termbin/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter 设置路由
func InitRouter(r *gin.Engine) {

	v1 := r.Group("api/v1")
	{
		v1.POST("user/sign-in", api.UserSignIn())
		v1.POST("user/sign-up", api.UserSignUp())

		authed := v1.Group("/")
		authed.Use(middleware.JWTAuth())
		{
			authed.POST("/", api.NewClipboard())
			authed.POST("/:alias", api.NewClipboard())
			authed.GET("/:id", api.GetClipboard())
			authed.PUT("/:id", api.UpdateClipboard())
			authed.DELETE("/:id", api.DeleteClipboard())
			authed.POST("/auth/:id", api.AuthorizeClipboard())
		}
	}
}
