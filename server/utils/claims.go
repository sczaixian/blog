package utils

import (
	"blog/server/global"
	"blog/server/models"
	"blog/server/models/request"
	"fmt"
	"net"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginToken(user models.Login) (token string, claims request.CustomClaims, err error) {
	j := NewJWT()
	claims = j.CreateClaims(request.BaseClaims{
		ID:       user.GetUserId(),
		Username: user.GetUserName(),
	})
	token, err = j.CreateToken(claims)
	return
}

func SetToken(c *gin.Context, token string, maxAge int) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", token, maxAge, "/", "", false, false)
	} else {
		c.SetCookie("x-token", token, maxAge, "/", host, false, false)
	}
}

func GetToken(c *gin.Context) string {
	token := c.Request.Header.Get("x-token")
	fmt.Println("----token--->> ", token)
	if token == "" {
		j := NewJWT()
		token, _ = c.Cookie("x-token")
		fmt.Println("---Cookie-token--->> ", token)
		claims, err := j.ParseToken(token)
		fmt.Println("---GetToken--->> ", claims, "-----", err)
		if err != nil {
			global.GVA_LOG.Error("重新写入cookie token失败,未能成功解析token,请检查请求头是否存在x-token且claims是否为规定结构")
			return token
		}
		SetToken(c, token, int((claims.ExpiresAt.Unix()-time.Now().Unix())/60))
	}
	return token
}

func GetClaims(c *gin.Context) (*request.CustomClaims, error) {
	fmt.Println("==---------GetClaims------------")
	token := GetToken(c)
	fmt.Println("----------GetClaims-----------------", token)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.BaseClaims.ID
	}
}

func ClearToken(c *gin.Context) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", "", -1, "/", "", false, false)
	} else {
		c.SetCookie("x-token", "", -1, "/", host, false, false)
	}
}
