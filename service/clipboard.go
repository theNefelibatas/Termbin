package service

import (
	"Termbin/dao"
	"Termbin/model"
	"Termbin/util"
	"io/ioutil"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// ClipboardSrv 剪切板服务
type ClipboardSrv struct {
}

// ClipboardSrvIns 剪切板服务单例
var ClipboardSrvIns *ClipboardSrv

// ClipboardSrvOnce 剪切板服务单例初始化锁
var ClipboardSrvOnce sync.Once

// GetClipboardSrv 获取剪切板服务单例
func GetClipboardSrv() *ClipboardSrv {
	ClipboardSrvOnce.Do(func() {
		ClipboardSrvIns = &ClipboardSrv{}
	})
	return ClipboardSrvIns
}

// NewClipboard 新建剪切板
func (s ClipboardSrv) NewClipboard(ctx *gin.Context, req *model.NewClipboardReq) (model.Clipboard, error) {
	// 获取剪贴板内容
	// content := ctx.PostForm("c")
	file, _ := ctx.FormFile("c")
	src, _ := file.Open()
	defer src.Close()

	content, _ := ioutil.ReadAll(src)

	// 生成剪贴板的哈希值
	digest, _ := util.GenDigest(string(content))

	// 生成剪贴板的 uuid
	uuid, _ := util.GenUUID()

	// 生成剪切板的 short
	short, _ := util.GenShortID(4)

	// 构造剪贴板对象
	clipboard := model.Clipboard{
		Date:    time.Now().Format(time.RFC3339),
		Digest:  digest,
		Short:   short,
		Size:    len(content),
		URL:     "http://127.0.0.1/" + short,
		Status:  "created",
		UUID:    uuid,
		Content: string(content),
	}

	// 在 DAO 层创建剪切板记录
	err := dao.NewClipboardDAO(ctx).NewClipboard(&clipboard)
	if err != nil {
		return model.Clipboard{}, err
	}

	return clipboard, nil
}

// GetClipboard 获取剪切板内容
func (s ClipboardSrv) GetClipboard(ctx *gin.Context, req *model.GetClipboardReq) (string, error) {
	// 获取剪贴板 ID
	id := ctx.Param("id")
	// id := req.ID

	// 在 DAO 层中查询剪贴板内容
	clipboard, err := dao.NewClipboardDAO(ctx).GetClipboard(id)
	if err != nil {
		return "", err
	}

	return clipboard.Content, nil
}

// UpdateClipboard 更新剪切板内容
func (s ClipboardSrv) UpdateClipboard(ctx *gin.Context, req *model.UpdateClipboardReq) (string, error) {
	// 获取剪贴板 ID
	id := ctx.Param("id")
	// id := req.ID

	// 获取剪贴板内容
	// content := ctx.PostForm("c")
	// content := req.Content
	file, _ := ctx.FormFile("c")
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	content, _ := ioutil.ReadAll(src)

	// 在 DAO 层中更新剪贴板内容
	clipboard, err := dao.NewClipboardDAO(ctx).UpdateClipboard(id, string(content))
	if err != nil {
		return "", err
	}

	return clipboard.URL, nil
}

// DeleteClipboard 删除剪切板
func (s ClipboardSrv) DeleteClipboard(ctx *gin.Context, req *model.DeleteClipboardReq) (string, error) {
	// 获取剪贴板 ID
	id := ctx.Param("id")
	// id := req.ID

	//clipboard, err := dao.NewClipboardDAO(ctx).GetClipboard(id)
	//if err != nil {
	//	return "", err
	//}

	// 在 DAO 层中删除剪贴板内容
	clipboard, err := dao.NewClipboardDAO(ctx).DeleteClipboard(id)
	if err != nil {
		return "", err
	}

	return clipboard.UUID, nil
}
