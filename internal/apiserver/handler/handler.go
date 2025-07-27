package handler

import "github.com/gin-gonic/gin"

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	apiv1 := server.Group("/api/v1")
	apiv1.POST("/users/signup", u.SignUp)
}
