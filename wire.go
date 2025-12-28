//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jwcen/mars/internal/apiserver/handler"
	ijwt "github.com/jwcen/mars/internal/apiserver/handler/jwt"
	"github.com/jwcen/mars/internal/apiserver/repository"
	"github.com/jwcen/mars/internal/apiserver/repository/cache"
	"github.com/jwcen/mars/internal/apiserver/repository/dao"
	"github.com/jwcen/mars/internal/apiserver/service"
	"github.com/jwcen/mars/ioc"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		ioc.InitGinEngine,
		ioc.InitMiddlewares,

		ioc.InitDB,
		ioc.InitRedis,

		dao.NewUserDao,

		cache.NewRedisUserCache,

		repository.NewUserInfoRepository,

		service.NewUserService,

		ijwt.NewRedisJWTHandler,
		handler.NewUserHandler,
	)

	return new(gin.Engine)
}
