package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserEmail string `gorm:"unique"`
	UserName  string
	Password  string
}

const passwordCost = 16

// SetPassword 设置密码并加密
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 检验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// UserSignUpReq 用户注册请求
type UserSignUpReq struct {
	UserEmail string `form:"user_email" json:"user_email" binding:"required"`
	UserName  string `form:"user_name" json:"user_name" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

// UserSignInReq 用户登录请求
type UserSignInReq struct {
	UserEmail string `form:"user_email" json:"user_email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

// UserResp 用户响应数据
type UserResp struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	Token     string `json:"token"`
}
