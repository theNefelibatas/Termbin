package model

import "gorm.io/gorm"

type Clipboard struct {
	gorm.Model
	Date    string `json:"date"`
	Digest  string `json:"digest"`
	Short   string `json:"short" gorm:"uniqueIndex"`
	Size    int    `json:"size"`
	URL     string `json:"url"`
	Status  string `json:"status"`
	UUID    string `json:"uuid" gorm:"uniqueIndex"`
	Content string `json:"content"`
}

type NewClipboardReq struct {
	Content string `form:"c" json:"content" binding:"required"`
}

type GetClipboardReq struct {
	ID string `json:"id"`
}

type UpdateClipboardReq struct {
	ID      string `json:"id"`
	Content string `form:"c" json:"content"`
}

type DeleteClipboardReq struct {
	ID string `json:"id"`
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
