package api

import (
	"Termbin/model"
	"Termbin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewClipboard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.NewClipboardReq
		//if err := ctx.ShouldBind(&req); err != nil {
		//	ctx.JSON(http.StatusBadRequest, gin.H{
		//		"err": "bind error",
		//	})
		//	return
		//}

		srv := service.GetClipboardSrv()
		clipboard, _ := srv.NewClipboard(ctx, &req)

		// c.JSON(http.StatusOK, clipboard)
		ctx.String(http.StatusOK, "date: %s\ndigest: %s\nshort: %s\nsize: %d\nurl: %s\nstatus: %s\nuuid: %s\n",
			clipboard.Date, clipboard.Digest, clipboard.Short, clipboard.Size, clipboard.URL, clipboard.Status, clipboard.UUID)
	}
}

func GetClipboard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.GetClipboardReq
		//if err := ctx.ShouldBind(&req); err != nil {
		//	ctx.JSON(http.StatusBadRequest, gin.H{})
		//	return
		//}

		srv := service.GetClipboardSrv()
		content, _ := srv.GetClipboard(ctx, &req)

		ctx.String(http.StatusOK, "%s\n", content)
	}
}

func UpdateClipboard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.UpdateClipboardReq
		//if err := ctx.ShouldBind(&req); err != nil {
		//	ctx.JSON(http.StatusBadRequest, gin.H{})
		//	return
		//}

		srv := service.GetClipboardSrv()
		url, _ := srv.UpdateClipboard(ctx, &req)

		ctx.String(http.StatusOK, "%s updated\n", url)
	}
}

func DeleteClipboard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.DeleteClipboardReq
		//if err := ctx.ShouldBind(&req); err != nil {
		//	ctx.JSON(http.StatusBadRequest, gin.H{})
		//	return
		//}

		srv := service.GetClipboardSrv()
		uuid, _ := srv.DeleteClipboard(ctx, &req)

		ctx.String(http.StatusOK, "deleted %s\n", uuid)
	}
}
