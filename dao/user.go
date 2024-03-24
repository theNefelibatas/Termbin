package dao

import (
	"Termbin/model"
	"context"

	"gorm.io/gorm"
)

type UserDAO struct {
	*gorm.DB
}

// NewUserDAO 创建 UserDAO 实例
func NewUserDAO(ctx context.Context) *UserDAO {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDAO{NewDBClient(ctx)}
}

// CreateUser 创建User
func (dao *UserDAO) CreateUser(user *model.User) error {
	err := dao.DB.Model(&model.User{}).Create(user).Error
	return err
}

// GetUserByUserName 在 user 表中根据用户名找到用户
func (dao *UserDAO) GetUserByUserName(userName string) (*model.User, error) {
	user := &model.User{}
	err := dao.DB.Model(&model.User{}).Where("user_name=?", userName).
		First(&user).Error

	return user, err
}

// GetUserByUserID 在 user 表中根据用户 id 找到用户
func (dao *UserDAO) GetUserByUserID(id uint) (*model.User, error) {
	user := &model.User{}
	err := dao.DB.Model(&model.User{}).Where("id=?", id).
		First(&user).Error

	return user, err
}
