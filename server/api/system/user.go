package system

import (
	"blog/server/global"
	"blog/server/models"
	common_response "blog/server/models/common/response"
	"blog/server/models/request"
	"blog/server/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (b *BaseApi) Login(c *gin.Context) {
	openCaptchaTimeOut := 60 // 缓存超时时间
	var login request.Login
	err := c.ShouldBind(&login)
	key := c.ClientIP()
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(login, utils.LoginVerify)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	fmt.Println(v)
	var oc = interfaceToInt(v)
	fmt.Println(oc)
	u := &models.User{UserName: login.Username, Password: login.Password}
	user, err := userService.Login(u)
	if err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		global.BlackCache.Increment(key, 1)
		common_response.FailWithMessage("用户名不存在或者密码错误", c)
		return
	}
	// TODO:多次登录错误
	b.TokenNext(c, *user)
	return
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}

// TokenNext 登录以后签发jwt
func (b *BaseApi) TokenNext(c *gin.Context, user models.User) {

}
