package user_web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ink-yht/hure/internal/domain/user_domain"
	"github.com/ink-yht/hure/internal/service/user_service"
	"github.com/ink-yht/hure/internal/web"
	"github.com/ink-yht/hure/pkg/logger"
	"net/http"
	"strconv"
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
	ug := server.Group("/users")
	ug.POST("/login", u.Login) // 接受登录凭证code和用户信息，返回 token
	ug.POST("/edit", u.Edit)
}

func (u *UserHandler) Login(ctx *gin.Context) {
	var req user_domain.CodeAndUserInfoRequest
	err := ctx.Bind(&req)
	if err != nil {
		// 出错会返回 400 错误
		return
	}
	fmt.Println("req", req)
	// 提前从请求头中提取 User-Agent 信息，用于后续可能的日志记录或审计。
	userAgent := ctx.GetHeader("User-Agent")

	token, err := u.svc.Login(ctx, req, userAgent)
	if err != nil {
		// 当出现其他错误时，返回通用的系统错误信息。
		ctx.JSON(http.StatusOK, web.Result{
			Code: 2,
			Msg:  "系统错误",
			Data: nil,
		})
		// 记录错误日志，包括错误详情。
		u.l.Error("系统错误", logger.Error("err", err))
		return
	}
	// 登录成功时，在响应头中设置JWT令牌，并返回成功信息。
	ctx.Header("x-jwt-token", token)
	ctx.JSON(http.StatusOK, web.Result{
		Code: 0,
		Msg:  "登录成功",
		Data: nil,
	})
	// 记录登录成功的日志，包括用户的邮箱信息。
	u.l.Info("登录成功", logger.String("nickname", req.NickName))
}

func (u *UserHandler) Edit(ctx *gin.Context) {
	userClaims := ctx.MustGet("claims").(*user_service.UserClaims)
	var req user_domain.RoleRequest
	err := ctx.Bind(&req)
	if err != nil {
		// 出错会返回 400 错误
		return
	}
	req.ID = userClaims.Id

	err = u.svc.Edit(ctx, req)
	if err != nil {
		// 当出现其他错误时，返回通用的系统错误信息。
		ctx.JSON(http.StatusOK, web.Result{
			Code: 2,
			Msg:  "系统错误",
			Data: nil,
		})
		// 记录错误日志，包括错误详情。
		u.l.Error("系统错误", logger.Error("err", err))
		return
	}
	ctx.JSON(http.StatusOK, web.Result{
		Code: 0,
		Msg:  "修改成功",
		Data: nil,
	})
	u.l.Info("登录成功", logger.String("ID", strconv.Itoa(int(req.ID))))
}
