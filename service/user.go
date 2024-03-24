package service

import (
	"Termbin/dao"
	"Termbin/model"
	"errors"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserSrv 用户服务
type UserSrv struct {
}

// UserSrvIns 用户服务单例
var UserSrvIns *UserSrv

// UserSrvOnce 用户服务单例初始化锁
var UserSrvOnce sync.Once

// GetUserSrv 获取用户服务单例
func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (s UserSrv) SignUp(ctx *gin.Context, req *model.UserReq) error {
	userDAO := dao.NewUserDAO(ctx)
	user, err := userDAO.GetUserByUserName(req.UserName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = &model.User{
			UserName: req.UserName,
		}
		// 密码加密存储
		if err = user.SetPassword(req.Password); err != nil {
			return err
		}
		if err = userDAO.CreateUser(user); err != nil {
			return err
		}
		return nil
	}
	if err == nil {
		return errors.New("user already exists")
	}
	return err
}

func (s UserSrv) SignIn(ctx *gin.Context, req *model.UserReq) (*model.UserResp, error) {
	userDAO := dao.NewUserDAO(ctx)
	user, err := userDAO.GetUserByUserName(req.UserName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user non-existent")
	}
	if !user.CheckPassword(req.Password) {
		return nil, errors.New("invalid username or password")
	}

	u := &model.UserResp{
		ID:       user.ID,
		UserName: user.UserName,
	}

	return u, nil
}
