package system

import (
	"blog/server/models"
)

type MenuService struct{}

var MenuServiceApp = new(MenuService)

func (MenuService *MenuService) UserAuthorityDefaultRouter(user *models.User) {
	//var menuIds []string
	//err := global.GVA_DB.Model(&models.)
	// todo:
}
