package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jwcen/mars/config"
	"github.com/jwcen/mars/internal/apiserver/handler"
	ijwt "github.com/jwcen/mars/internal/apiserver/handler/jwt"
	"github.com/jwcen/mars/internal/apiserver/repository"
	"github.com/jwcen/mars/internal/apiserver/repository/dao"
	"github.com/jwcen/mars/internal/apiserver/service"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()
	r := initWebServer()
	u := initUser(db)
	u.RegisterRoutes(r)
	err := r.Run(":8081")
	if err != nil {
		panic("端口启动失败")
	}
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initWebServer() *gin.Engine {
	r := gin.Default()
	return r
}

func initUser(db *gorm.DB) *handler.UserHandler {
	da := dao.NewUserDao(db)
	repo := repository.NewUserInfoRepository(da)
	svc := service.NewUserService(repo)
	cmd := redis.NewClient(&redis.Options{
		Addr: config.Config.Redis.Addr,
	})
	jwtHandler := ijwt.NewRedisJWTHandler(cmd)
	u := handler.NewUserHandler(svc, jwtHandler)
	return u
}
