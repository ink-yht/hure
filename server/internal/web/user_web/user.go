package user_web

import (
	"github.com/gin-gonic/gin"
	"github.com/ink-yht/hure/internal/service/user_service"
	"github.com/ink-yht/hure/internal/web"
	"github.com/ink-yht/hure/pkg/logger"
)

// 确保 UserHandler 上实现了 Handler 接口
var _ web.Handler = (*UserHandler)(nil)

type UserHandler struct {
	svc user_service.UserService
	l   logger.Logger
}

func NewUserHandler(svc user_service.UserService, l logger.Logger) *UserHandler {
	return &UserHandler{
		svc: svc,
		l:   l,
	}
}

// RegisterRoutes 路由注册
func (u *UserHandler) RegisterRoutes(server *gin.Engine) {

}
