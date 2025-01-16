//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/ink-yht/hure/internal/repository/dao/user_dao"
	"github.com/ink-yht/hure/internal/repository/user_repo"
	"github.com/ink-yht/hure/internal/service/user_service"
	"github.com/ink-yht/hure/internal/web/user_web"
	"github.com/ink-yht/hure/ioc"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 最基础的第三方依赖
		ioc.InitDB, ioc.InitLogger,

		// DAO 部分
		user_dao.NewUserDAO,
		user_dao.NewAdminDAO,

		// cache 部分

		// repository 部分
		user_repo.NewUserRepository,
		user_repo.NewAdminRepository,

		// service 部分
		user_service.NewUserService,
		user_service.NewAdminService,

		// Handler 部分
		user_web.NewUserHandler,
		user_web.NewAdminHandler,

		// 中间件
		ioc.InitWebServer,
		ioc.InitMiddleWares,
	)
	return new(gin.Engine)
}
