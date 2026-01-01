package ioc

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jwcen/mars/internal/apiserver/handler"
	"github.com/jwcen/mars/internal/apiserver/handler/middleware"
	"github.com/jwcen/mars/pkg/ginx/middlewares/ratelimit"
	"github.com/redis/go-redis/v9"
)

func InitGinEngine(mdls []gin.HandlerFunc,
	userHdl *handler.UserHandler,
	articleHdl *handler.ArticleHandler) *gin.Engine {
	
	server := gin.Default()
	server.Use(gin.Recovery())
	server.Use(gin.Logger())
	server.Use(gin.ErrorLogger())
	server.Use(mdls...)

	userHdl.RegisterRoutes(server)
	articleHdl.RegisterRoutes(server)
	return server
}

func InitMiddlewares(redisClient redis.Cmdable) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		corsHdl(),
		middleware.NewLoginJWTMiddlewareBuilder().
			IgnorePaths("/api/v1/users/signup").
			IgnorePaths("/api/v1/users/login_sms/code/send").
			IgnorePaths("/api/v1/users/login_sms").
			IgnorePaths("/api/v1/users/login").Build(),
		ratelimit.NewBuilder(redisClient, time.Second, 100).Build(),
	}
}

func corsHdl() gin.HandlerFunc {
	return cors.New(cors.Config{
		//AllowOrigins: []string{"*"},
		//AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		// 你不加这个，前端是拿不到的
		ExposeHeaders: []string{"x-jwt-token"},
		// 是否允许你带 cookie 之类的东西
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				// 你的开发环境
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	})
}
