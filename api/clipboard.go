package api

import (
	"Termbin/model"
	"Termbin/service"
	"Termbin/util"
	"net/http"
	"strings"

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
		req.ID, _ = util.RemoveExt(ctx.Param("id"))
		srv := service.GetClipboardSrv()
		clipboard, err := srv.NewClipboard(ctx, &req)
		if err != nil {
			ctx.String(http.StatusOK, "date:%s\nsize:%s\nstatus:%s\n", clipboard.Date, clipboard.Size, clipboard.Status)
			return
		}

		userAgent := ctx.GetHeader("User-Agent")
		if strings.Contains(userAgent, "curl") {
			ctx.String(http.StatusOK, "date: %s\ndigest: %s\nshort: %s\nsize: %d\nurl: %s\nstatus: %s\nuuid: %s\n",
				clipboard.Date, clipboard.Digest, clipboard.Short, clipboard.Size, clipboard.URL, clipboard.Status, clipboard.UUID)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data": clipboard,
			})
		}
	}
}

func GetClipboard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := model.ClipboardReq{
			ID: ctx.Param("id"),
		}
		req.ID, _ = util.RemoveExt(req.ID)

		srv := service.GetClipboardSrv()
		content, err := srv.GetClipboard(ctx, &req)
		if err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}

		userAgent := ctx.GetHeader("User-Agent")
		if strings.Contains(userAgent, "curl") {
			ctx.String(http.StatusOK, "%s\n", content)
		} else {
			contentType, _ := util.GetContentType(ctx.Request.URL.Path)
			if contentType != "" {
				// ctx.Header("content-type", contentType)
				ctx.Data(http.StatusOK, contentType, []byte(content))
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"data": content,
				})
			}

		}
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
		req.ID, _ = util.RemoveExt(ctx.Param("id"))

		srv := service.GetClipboardSrv()
		url, err := srv.UpdateClipboard(ctx, &req)
		if err != nil {
			ctx.String(http.StatusOK, "failed to update %s: "+err.Error(), url)
			return
		}

		userAgent := ctx.GetHeader("User-Agent")
		if strings.Contains(userAgent, "curl") {
			ctx.String(http.StatusOK, "%s updated\n", url)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data": url,
			})
		}

	}
}

func DeleteClipboard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := model.ClipboardReq{
			ID: ctx.Param("id"),
		}
		req.ID, _ = util.RemoveExt(req.ID)

		srv := service.GetClipboardSrv()
		uuid, err := srv.DeleteClipboard(ctx, &req)
		if err != nil {
			ctx.String(http.StatusOK, "failed to delete %s: "+err.Error(), uuid)
			return
		}

		userAgent := ctx.GetHeader("User-Agent")
		if strings.Contains(userAgent, "curl") {
			ctx.String(http.StatusOK, "deleted %s\n", uuid)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data": uuid,
			})
		}

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

		userAgent := ctx.GetHeader("User-Agent")
		if strings.Contains(userAgent, "curl") {
			ctx.String(http.StatusOK, "authorize %s to %s", req.UserEmail, url)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data": gin.H{
					"user_email": req.UserEmail,
					"url":        url,
				},
			})
		}
	}
}
