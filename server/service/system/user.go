package system

import (
	"blog/server/global"
	"blog/server/models"
	"blog/server/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
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

func (service *UserService) Register(u models.User) (userInit models.User, err error) {
	var user models.User
	if !errors.Is(global.GVA_DB.Where("username = ?", u.UserName).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInit, errors.New("用户名已注册")
	}
	u.Password = utils.BcryptHash(u.Password)
	err = global.GVA_DB.Create(&u).Error
	return u, err
}

func (service *UserService) Logout() {

}

func (service *UserService) ResetPassword(ID uint, password string) (err error) {
	err = global.GVA_DB.Model(&models.User{}).Where("id = ?", ID).Update("password", utils.BcryptHash(password)).Error
	return err
}

func (service *UserService) ChangePassword(u *models.User, newPassword string) (err error) {
	var user models.User
	err = global.GVA_DB.Select("id, password").Where("id = ?", u.ID).First(&user).Error
	if err != nil {
		return err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return errors.New("原密码错误")
	}
	pwd := utils.BcryptHash(newPassword)
	err = global.GVA_DB.Model(&user).Update("password", pwd).Error
	return err
}
