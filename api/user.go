package api

import (
	"Termbin/model"
	"Termbin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.UserReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": "bind error",
			})
			return
		}
		srv := service.GetUserSrv()
		err := srv.SignUp(ctx, &req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": "sign up error",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})

	}
}

func UserSignIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.UserReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": "bind error",
			})
			return
		}
		srv := service.GetUserSrv()
		err := srv.SignIn(ctx, &req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": "sign in error",
			})
			return
		}
	}
}
