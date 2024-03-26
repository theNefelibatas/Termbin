package api

import (
	"Termbin/model"
	"Termbin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewClipboard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.ClipboardReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": "bind error",
			})
			return
		}
		req.ID = ctx.Param("id")
		srv := service.GetClipboardSrv()
		clipboard, err := srv.NewClipboard(ctx, &req)
		if err != nil {
			ctx.String(http.StatusOK, "date:%s\nsize:%s\nstatus:%s\n", clipboard.Date, clipboard.Size, clipboard.Status)
			return
		}

		//ctx.JSON(http.StatusOK, gin.H{
		//	"data": clipboard,
		//})
		ctx.String(http.StatusOK, "date: %s\ndigest: %s\nshort: %s\nsize: %d\nurl: %s\nstatus: %s\nuuid: %s\n",
			clipboard.Date, clipboard.Digest, clipboard.Short, clipboard.Size, clipboard.URL, clipboard.Status, clipboard.UUID)
	}
}

func GetClipboard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := model.ClipboardReq{
			ID: ctx.Param("id"),
		}

		srv := service.GetClipboardSrv()
		content, err := srv.GetClipboard(ctx, &req)
		if err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}

		ctx.String(http.StatusOK, "%s\n", content)
	}
}

func UpdateClipboard() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req model.ClipboardReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": "bind error",
			})
			return
		}
		req.ID = ctx.Param("id")

		srv := service.GetClipboardSrv()
		url, err := srv.UpdateClipboard(ctx, &req)
		if err != nil {
			ctx.String(http.StatusOK, "failed to update %s: "+err.Error(), url)
			return
		}

		ctx.String(http.StatusOK, "%s updated\n", url)
	}
}

func DeleteClipboard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := model.ClipboardReq{
			ID: ctx.Param("id"),
		}

		srv := service.GetClipboardSrv()
		uuid, err := srv.DeleteClipboard(ctx, &req)
		if err != nil {
			ctx.String(http.StatusOK, "failed to delete %s: "+err.Error(), uuid)
			return
		}

		ctx.String(http.StatusOK, "deleted %s\n", uuid)
	}
}

func AuthorizeClipboard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.AuthClipboardReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": "bind error",
			})
			return
		}
		req.ID = ctx.Param("id")

		srv := service.GetClipboardSrv()
		url, err := srv.AuthorizeClipboard(ctx, &req)
		if err != nil {
			ctx.String(http.StatusOK, "failed to authorize %s: "+err.Error(), req.UserEmail)
			return
		}

		if req.UserEmail == "" {
			ctx.String(http.StatusOK, "only author has access to %s", url)
			return
		}
		ctx.String(http.StatusOK, "authorize %s to %s", req.UserEmail, url)
	}
}
