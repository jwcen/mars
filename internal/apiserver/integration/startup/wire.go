//go:build wireinject
// +build wireinject

package startup

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

var thirdProvider = wire.NewSet(InitRedis, InitTestDB)
var userSvcProvider = wire.NewSet(
	dao.NewUserDao,
	cache.NewRedisUserCache,
	repository.NewUserInfoRepository,
	service.NewUserService)

func InitWebServer() *gin.Engine {
	wire.Build(
		ioc.InitGinEngine,
		ioc.InitMiddlewares,

		thirdProvider,
		userSvcProvider,

		dao.NewArticleDao,
		repository.NewArticleRepository,
		service.NewArticleService,

		ijwt.NewRedisJWTHandler,
		handler.NewUserHandler,
		handler.NewArticleHandler,
	)

	return new(gin.Engine)
}

func InitArticleHandler() *handler.ArticleHandler {
	wire.Build(
		thirdProvider,
		dao.NewArticleDao,
		repository.NewArticleRepository,
		service.NewArticleService,
		handler.NewArticleHandler,
	)
	return &handler.ArticleHandler{}
}
