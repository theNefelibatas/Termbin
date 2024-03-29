package api

import (
	"Termbin/model"
	"Termbin/service"
	"Termbin/util"
	"bytes"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/gin-gonic/gin"
)

func NewClipboard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.ClipboardReq
		//if err := ctx.ShouldBind(&req); err != nil {
		//	ctx.JSON(http.StatusBadRequest, gin.H{
		//		"err": "bind error",
		//	})
		//	return
		//}
		file, err := ctx.FormFile("content")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": "invalid file:" + err.Error(),
			})
			return
		}
		content, _ := file.Open()
		defer content.Close()
		req.Content, _ = ioutil.ReadAll(content)

		sunset := ctx.PostForm("sunset")
		if sunset != "" {
			req.Sunset, err = strconv.Atoi(sunset)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"err": "invalid sunset value",
				})
				return
			}
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
		codeFileExt := model.CodeFileExt[filepath.Ext(ctx.Request.URL.Path)]
		if strings.Contains(userAgent, "curl") {
			ctx.String(http.StatusOK, "%s", content)
		} else if util.IsBrowserUserAgent(userAgent) && codeFileExt != "" {

			lexer := lexers.Get(codeFileExt)
			if lexer == nil {
				lexer = lexers.Fallback
			}
			lexer = chroma.Coalesce(lexer)
			style := styles.Get("github")
			if style == nil {
				style = styles.Fallback
			}
			formatter := formatters.Get("html")
			if formatter == nil {
				formatter = formatters.Fallback
			}
			iterator, _ := lexer.Tokenise(nil, string(content))
			var buf bytes.Buffer

			_ = formatter.Format(&buf, style, iterator)
			result := buf.String()

			ctx.Header("Content-Type", "text/html")
			ctx.String(http.StatusOK, "%s", result)
		} else {
			contentType, _ := util.GetContentType(ctx.Request.URL.Path)
			if contentType != "" {
				// ctx.Header("content-type", contentType)
				ctx.Data(http.StatusOK, contentType, content)
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"data": string(content),
				})
			}
		}
	}
}

func UpdateClipboard() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req model.ClipboardReq
		//if err := ctx.ShouldBind(&req); err != nil {
		//	ctx.JSON(http.StatusBadRequest, gin.H{
		//		"err": "bind error",
		//	})
		//	return
		//}

		file, err := ctx.FormFile("content")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": "invalid file:" + err.Error(),
			})
			return
		}
		content, _ := file.Open()
		defer content.Close()
		req.Content, _ = ioutil.ReadAll(content)

		sunset := ctx.PostForm("sunset")
		if sunset != "" {
			req.Sunset, err = strconv.Atoi(sunset)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"err": "invalid sunset value",
				})
				return
			}
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
