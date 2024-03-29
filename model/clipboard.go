package model

import (
	"gorm.io/gorm"
)

type Clipboard struct {
	gorm.Model
	Author       *uint
	AllowedUsers *uint
	Access       int
	Date         string
	Digest       string
	Short        string  `gorm:"uniqueIndex"`
	Alias        *string `gorm:"unique"`
	URL          string
	UUID         string `gorm:"uniqueIndex"`
	Size         int
	//Content      string
	Content []byte `gorm:"type:longblob"`
	Burn    bool
}

type ClipboardReq struct {
	ID string `form:"id" json:"id"`
	//Content string `form:"content" json:"content"`
	Content []byte `form:"content" json:"content"`
	Sunset  int    `form:"sunset" json:"sunset"`
}

type AuthClipboardReq struct {
	ID        string `form:"id" json:"id"`
	UserEmail string `form:"user_email" json:"user_email"`
	Burn      bool   `form:"burn" json:"burn"`
}

type ClipboardResp struct {
	Date   string `json:"date"`
	Digest string `json:"digest"`
	Short  string `json:"short"`
	Size   int    `json:"size"`
	URL    string `json:"url"`
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	//Content string `json:"content"`
	Content []byte `json:"content"`
}
