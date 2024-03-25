package service

import (
	"Termbin/dao"
	"Termbin/model"
	"Termbin/util"
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

func (s UserSrv) SignUp(ctx *gin.Context, req *model.UserSignUpReq) error {
	userDAO := dao.NewUserDAO(ctx)
	user, err := userDAO.GetUserByUserEmail(req.UserEmail)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = &model.User{
			UserEmail: req.UserEmail,
			UserName:  req.UserName,
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

func (s UserSrv) SignIn(ctx *gin.Context, req *model.UserSignInReq) (*model.UserResp, error) {
	userDAO := dao.NewUserDAO(ctx)
	user, err := userDAO.GetUserByUserEmail(req.UserEmail)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user non-existent")
	}
	if !user.CheckPassword(req.Password) {
		return nil, errors.New("invalid email or password")
	}
	token, err := util.GenerateToken(user.ID, user.UserEmail)
	if err != nil {

	}

	u := &model.UserResp{
		ID:        user.ID,
		UserName:  user.UserName,
		UserEmail: user.UserEmail,
		Token:     token,
	}

	return u, nil
}
