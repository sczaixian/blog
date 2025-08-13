package router

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user") //.Use(middleware)
	{
		userRouter.POST("login", UserApi.Login)
		userRouter.POST("register", UserApi.Register)
		userRouter.POST("changePassword", UserApi.ChangePassword)
		userRouter.POST("resetPassword", UserApi.ResetPassword)
	}
}
