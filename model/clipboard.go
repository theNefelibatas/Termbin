package model

import "gorm.io/gorm"

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
	Content      string
}

type ClipboardReq struct {
	ID      string `form:"id" json:"id"`
	Content string `form:"content" json:"content"`
}

type ClipboardResp struct {
	Date    string `json:"date"`
	Digest  string `json:"digest"`
	Short   string `json:"short"`
	Size    int    `json:"size"`
	URL     string `json:"url"`
	Status  string `json:"status"`
	UUID    string `json:"uuid"`
	Content string `json:"content"`
}
