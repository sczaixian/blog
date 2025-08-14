package system

import (
	"blog/server/global"
	"blog/server/models"
)

type CategoryService struct{}

func (cs *CategoryService) CreateCategory(category *models.Category) error {
	err := global.GVA_DB.Create(category).Error
	return err
}
