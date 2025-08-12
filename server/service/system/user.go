package system

import (
	"blog/server/global"
	"blog/server/models"
	"blog/server/utils"
	"fmt"
)

type UserService struct{}

func (service *UserService) Login(u *models.User) (userInit *models.User, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}
	var user models.User
	err = global.GVA_DB.Where("username = ?", u.UserName).Preload("Articles").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, fmt.Errorf("password error")
		}
		return nil, err
	}
	return u, nil
}
