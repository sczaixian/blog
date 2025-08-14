package system

import (
	"blog/server/global"
	"blog/server/models"
	"context"
)

type JwtService struct{}

var JwtServiceApp = new(JwtService)

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.GVA_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

func (jwtService *JwtService) JsonInBlacklist(jwtList models.Jwt) (err error) {
	err = global.GVA_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	//global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//
//func SetRedisJWT(jwt string, userName string) (err error) {
//	// 此处过期时间等于jwt过期时间
//	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
//	if err != nil {
//		return err
//	}
//	timer := dr
//	err = global.GVA_REDIS.Set(context.Background(), userName, jwt, timer).Err()
//	return err
//}
