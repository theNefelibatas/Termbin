package dao

import (
	"Termbin/model"
	"context"

	"gorm.io/gorm"
)

type ClipboardDAO struct {
	*gorm.DB
}

func NewClipboardDAO(ctx context.Context) *ClipboardDAO {
	if ctx == nil {
		ctx = context.Background()
	}
	return &ClipboardDAO{NewDBClient(ctx)}
}

func (dao *ClipboardDAO) NewClipboard(clipboard *model.Clipboard) error {
	return dao.Model(&model.Clipboard{}).Create(&clipboard).Error
}

func (dao *ClipboardDAO) GetClipboard(id string) (*model.Clipboard, error) {
	clipboard := &model.Clipboard{}
	if err := dao.Model(&model.Clipboard{}).Where("short = ? OR uuid = ?", id, id).First(&clipboard).Error; err != nil {
		return nil, err
	}
	return clipboard, nil
}

func (dao *ClipboardDAO) UpdateClipboard(id string, content string) (*model.Clipboard, error) {
	clipboard := &model.Clipboard{}
	if err := dao.Model(&model.Clipboard{}).Where("short = ? OR uuid = ?", id, id).First(&clipboard).Error; err != nil {
		return nil, err
	}
	clipboard.Content = content
	clipboard.Size = len(content)
	if err := dao.Model(&model.Clipboard{}).Where("short = ? OR uuid = ?", id, id).Updates(&clipboard).Error; err != nil {
		return nil, err
	}

	return clipboard, nil
}

func (dao *ClipboardDAO) DeleteClipboard(id string) (*model.Clipboard, error) {
	clipboard := &model.Clipboard{}
	if err := dao.Model(&model.Clipboard{}).Where("short = ? OR uuid = ?", id, id).First(&clipboard).Error; err != nil {
		return nil, err
	}
	if err := dao.Model(&model.Clipboard{}).Delete(&clipboard).Error; err != nil {
		return nil, err
	}
	return clipboard, nil
}
