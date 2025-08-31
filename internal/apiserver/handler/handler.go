package handler

import "github.com/gin-gonic/gin"

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	apiv1 := server.Group("/api/v1")
	ug := apiv1.Group("/users")
	{
		ug.POST("/signup", u.SignUp)
		ug.POST("/login", u.Login)
		ug.POST("/logout", u.Logout)
		ug.GET("/profile", u.Profile)

		ug.POST("/login_sms/code/send", u.SendLoginSMSCode)
		ug.POST("/login_sms", u.LoginSMS)
	}
}
