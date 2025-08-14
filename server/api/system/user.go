package system

import (
	"blog/server/global"
	"blog/server/models"
	common_response "blog/server/models/common/response"
	"blog/server/models/request"
	"blog/server/models/response"
	"blog/server/utils"

	"fmt"
	//"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct{}

func (u *UserApi) Login(c *gin.Context) {
	//openCaptchaTimeOut := 60 // 缓存超时时间
	var login request.Login
	err := c.ShouldBind(&login)
	key := c.ClientIP()
	fmt.Println(key)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(login, utils.LoginVerify)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}

	//查缓存
	//v, ok := global.BlackCache.Get(key)
	//if !ok {
	//	global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	//}
	//fmt.Println(v)
	//var oc = interfaceToInt(v)
	//fmt.Println(oc)

	us := &models.User{UserName: login.Username, Password: login.Password}
	user, err := userService.Login(us)
	if err != nil {
		fmt.Println("登陆失败! 用户名不存在或者密码错误!", err)
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		//global.BlackCache.Increment(key, 1)
		common_response.FailWithMessage("用户名不存在或者密码错误", c)
		return
	}
	// TODO:多次登录错误, TokenNext实现, 更新ip
	u.TokenNext(c, *user)
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
func (u *UserApi) TokenNext(c *gin.Context, user models.User) {
	//token, claims, err := utils.Lo
	common_response.OkWithDetailed(response.LoginResponse{
		User:      user,
		Token:     "token_123456",
		ExpiresAt: 200000 * 1000,
	}, "登录成功", c)
}

func (u *UserApi) Register(c *gin.Context) {
	var r request.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	user := &models.User{UserName: r.Username, Password: r.Password, Email: r.Email}
	userReturn, err := userService.Register(*user)

	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		common_response.FailWithDetailed(response.SysUserResponse{User: userReturn}, "注册失败", c)
		return
	}
	common_response.OkWithDetailed(response.SysUserResponse{User: userReturn}, "注册成功", c)
}

func (u *UserApi) ChangePassword(c *gin.Context) {
	var req request.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	us := &models.User{BaseModel: models.BaseModel{ID: uid}, Password: req.NewPassword}
	err = userService.ChangePassword(us, req.NewPassword)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		common_response.FailWithMessage("修改失败，原密码与当前账户不符", c)
		return
	}
	common_response.OkWithMessage("修改成功", c)
}

func (u *UserApi) ResetPassword(c *gin.Context) {
	var rps request.ResetPassword
	err := c.ShouldBindJSON(&rps)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.ResetPassword(rps.ID, rps.Password)
	if err != nil {
		global.GVA_LOG.Error("重置失败!", zap.Error(err))
		common_response.FailWithMessage("重置失败"+err.Error(), c)
		return
	}
	common_response.OkWithMessage("重置成功", c)
}
