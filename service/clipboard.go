package service

import (
	"Termbin/dao"
	"Termbin/model"
	"Termbin/util"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// ClipboardSrv 剪切板服务
type ClipboardSrv struct {
	mutex sync.Mutex
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
func (s *ClipboardSrv) NewClipboard(ctx *gin.Context, req *model.ClipboardReq) (model.ClipboardResp, error) {
	clipboardDAO := dao.NewClipboardDAO(ctx)

	// 获取剪贴板内容
	content := req.Content
	sunset := time.Duration(req.Sunset) * time.Second
	fmt.Println(sunset)

	// 生成剪贴板的哈希值
	digest, _ := util.GenDigest(content)

	// 生成剪贴板的 uuid
	uuid, _ := util.GenUUID()

	// 生成剪切板的 short
	short, _ := util.GenShortID(4)

	// 构造剪贴板对象
	clipboard := model.Clipboard{
		Author:       nil,
		AllowedUsers: nil,
		Access:       model.AllAccess,
		Date:         time.Now().Format(time.RFC3339),
		Digest:       digest,
		Short:        short,
		Alias:        nil,
		Size:         len(content),
		URL:          "http://127.0.0.1/api/v1/" + short,
		UUID:         uuid,
		Content:      content,
		Burn:         false,
	}

	if req.ID != "" {
		clipboard.Alias = &req.ID
	}

	// 检查登录状态
	userID, exist := ctx.Get("UserID")
	if exist {
		authorID := userID.(uint)
		clipboard.Author = &authorID
	}

	// 在 DAO 层创建剪切板记录
	err := clipboardDAO.NewClipboard(&clipboard)
	if err != nil {
		return model.ClipboardResp{
			Date:   clipboard.Date,
			Size:   clipboard.Size,
			Status: "failed",
		}, err
	}

	if sunset > 0 {
		time.AfterFunc(sunset, func() {
			s.mutex.Lock()
			defer s.mutex.Unlock()
			// 在超时后将权限设为 NoneAccess
			clipboard.Access = model.NoneAccess
			_ = clipboardDAO.UpdateClipboard(clipboard.Short, &clipboard) // 更新剪切板记录
		})
	}

	resp := model.ClipboardResp{
		Date:   clipboard.Date,
		Digest: clipboard.Digest,
		Short:  clipboard.Short,
		Size:   clipboard.Size,
		URL:    clipboard.URL,
		Status: "created",
		UUID:   clipboard.UUID,
	}

	return resp, nil
}

// GetClipboard 获取剪切板内容
func (s *ClipboardSrv) GetClipboard(ctx *gin.Context, req *model.ClipboardReq) ([]byte, error) {
	clipboardDAO := dao.NewClipboardDAO(ctx)

	// 获取剪贴板 ID
	id := req.ID
	// 在 DAO 层中查询剪贴板内容
	clipboard, err := clipboardDAO.GetClipboard(id)
	if err != nil {
		return nil, err
	}

	switch clipboard.Access {
	case model.AllAccess:
		return clipboard.Content, nil
	case model.AuthorAccess:
		userID, exist := ctx.Get("UserID")
		if !exist {
			return nil, errors.New("access denied")
		}
		if userID.(uint) != *clipboard.Author {
			return nil, errors.New("access denied")
		}
		return clipboard.Content, nil
	case model.AuthorizedAccess:
		userID, exist := ctx.Get("UserID")
		if !exist {
			return nil, errors.New("access denied")
		}
		if userID.(uint) == *clipboard.Author {
			return clipboard.Content, nil
		}
		if userID.(uint) == *clipboard.AllowedUsers {
			if clipboard.Burn {
				clipboard.Burn = false
				clipboard.Access = model.AuthorAccess
				clipboard.AllowedUsers = nil
				_ = clipboardDAO.UpdateClipboard(id, clipboard)
			}
			return clipboard.Content, nil
		}
		return nil, errors.New("access denied")
	case model.NoneAccess:
		return []byte("expired"), nil
	}
	return nil, errors.New("wtf")
}

// UpdateClipboard 更新剪切板内容
func (s *ClipboardSrv) UpdateClipboard(ctx *gin.Context, req *model.ClipboardReq) (string, error) {
	clipboardDAO := dao.NewClipboardDAO(ctx)
	// 获取剪贴板 ID
	id := req.ID
	// 获取剪贴板待更新内容
	content := req.Content

	clipboard, err := clipboardDAO.GetClipboard(id)
	if err != nil {
		return "", err
	}

	if clipboard.Access == model.NoneAccess {
		return clipboard.URL + "expired", nil
	}

	clipboard.Content = content
	clipboard.Size = len(content)
	clipboard.Digest, _ = util.GenDigest(content)

	if clipboard.Author == nil {
		err := clipboardDAO.UpdateClipboard(id, clipboard)
		return clipboard.URL, err
	}

	userID, exist := ctx.Get("UserID")
	if !exist {
		return clipboard.URL, errors.New("access denied")
	}
	if userID.(uint) != *clipboard.Author {
		return clipboard.URL, errors.New("access denied")
	}
	err = clipboardDAO.UpdateClipboard(id, clipboard)
	return clipboard.URL, err

	//switch clipboard.Access {
	//case model.AllAccess:
	//	err := clipboardDAO.UpdateClipboard(id, clipboard)
	//	return clipboard.URL, err
	//default:
	//	userID, exist := ctx.Get("UserID")
	//	if !exist {
	//		return clipboard.URL, errors.New("access denied")
	//	}
	//	if userID.(uint) != *clipboard.Author {
	//		return clipboard.URL, errors.New("access denied")
	//	}
	//	err := clipboardDAO.UpdateClipboard(id, clipboard)
	//	return clipboard.URL, err
	//}
}

// DeleteClipboard 删除剪切板
func (s *ClipboardSrv) DeleteClipboard(ctx *gin.Context, req *model.ClipboardReq) (string, error) {
	clipboardDAO := dao.NewClipboardDAO(ctx)

	// 获取剪贴板 ID
	id := req.ID

	clipboard, err := clipboardDAO.GetClipboard(id)
	if err != nil {
		return "", err
	}

	if clipboard.Author == nil {
		err := clipboardDAO.DeleteClipboard(id, clipboard)
		return clipboard.UUID, err
	}

	userID, exist := ctx.Get("UserID")
	if !exist {
		return clipboard.UUID, errors.New("access denied")
	}
	if userID.(uint) != *clipboard.Author {
		return clipboard.UUID, errors.New("access denied")
	}
	err = clipboardDAO.DeleteClipboard(id, clipboard)
	return clipboard.UUID, err

	//switch clipboard.Access {
	//case model.AllAccess:
	//	err := clipboardDAO.DeleteClipboard(id, clipboard)
	//	return clipboard.UUID, err
	//default:
	//	userID, exist := ctx.Get("UserID")
	//	if !exist {
	//		return clipboard.UUID, errors.New("access denied")
	//	}
	//	if userID.(uint) != *clipboard.Author {
	//		return clipboard.UUID, errors.New("access denied")
	//	}
	//	err := clipboardDAO.DeleteClipboard(id, clipboard)
	//	return clipboard.UUID, err
	//}
}

// AuthorizeClipboard 给剪切板设置指定用户可见
func (s *ClipboardSrv) AuthorizeClipboard(ctx *gin.Context, req *model.AuthClipboardReq) (string, error) {
	userDAO := dao.NewUserDAO(ctx)
	clipboardDAO := dao.NewClipboardDAO(ctx)

	id := req.ID
	userEmail := req.UserEmail
	burn := req.Burn

	clipboard, err := clipboardDAO.GetClipboard(id)
	if err != nil {
		return "", err
	}
	if clipboard.Access == model.NoneAccess {
		return clipboard.URL + "expired", nil
	}

	if clipboard.Author == nil {
		return clipboard.URL, errors.New("cannot authorize since the clipboard authorless")
	}
	if userEmail == "" {
		clipboard.Access = model.AuthorAccess
		if burn {
			return clipboard.URL, errors.New("cannot set read and burn when author read only")
		}
		// fmt.Println("empty user email")
	} else {
		// fmt.Println("user email is " + userEmail)
		user, err := userDAO.GetUserByUserEmail(userEmail)
		if err != nil {
			return clipboard.URL, errors.New("invalid user email")
		}
		clipboard.Access = model.AuthorizedAccess
		clipboard.AllowedUsers = &user.ID
		clipboard.Burn = burn
	}

	userID, exist := ctx.Get("UserID")
	if !exist {
		return clipboard.URL, errors.New("access denied")
	}
	if userID.(uint) != *clipboard.Author {
		return clipboard.URL, errors.New("access denied")
	}
	err = clipboardDAO.UpdateClipboard(id, clipboard)
	return clipboard.URL, err
}
