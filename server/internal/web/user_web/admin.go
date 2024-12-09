package user_web

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ink-yht/hure/internal/domain/user_domain"
	"github.com/ink-yht/hure/internal/service/user_service"
	"github.com/ink-yht/hure/internal/web"
	"github.com/ink-yht/hure/pkg/logger"
	"net/http"
)

// 确保 UserHandler 上实现了 Handler 接口
var _ web.Handler = (*AdminHandler)(nil)

type AdminHandler struct {
	svc user_service.AdminService
	l   logger.Logger
}

func NewAdminHandler(svc user_service.AdminService, l logger.Logger) *AdminHandler {
	return &AdminHandler{
		svc: svc,
		l:   l,
	}
}

// RegisterRoutes 路由注册
func (a *AdminHandler) RegisterRoutes(server *gin.Engine) {
	ag := server.Group("/admins")
	ag.POST("/signup", a.SignUp) // 管理员注册
}

// SignUp 管理员注册
func (a *AdminHandler) SignUp(ctx *gin.Context) {
	var req user_domain.AdminRegisterRequest
	err := ctx.Bind(&req)
	if err != nil {
		// 出错会返回 400 错误
		return
	}

	err = a.svc.Signup(ctx, req)

	if errors.Is(err, user_domain.ErrTheMailboxIsNotInTheRightFormat) {
		ctx.JSON(http.StatusOK, web.Result{
			Code: 1,
			Msg:  "电子邮件格式无效",
			Data: nil,
		})
		a.l.Warn("电子邮件格式无效", logger.String("email", req.Email))
		return
	}
	if errors.Is(err, user_domain.ErrThePasswordIsNotInTheRightFormat) {
		ctx.JSON(http.StatusOK, web.Result{
			Code: 1,
			Msg:  "密码长度必须为 8-20 个字符，并包含字母、数字和特殊字符",
			Data: nil,
		})
		a.l.Warn("密码格式不对", logger.String("email", req.Email))
		return
	}
	if errors.Is(err, user_domain.ErrThePasswordIsInconsistentTwice) {
		ctx.JSON(http.StatusOK, web.Result{
			Code: 1,
			Msg:  "两次密码不一致",
			Data: nil,
		})
		a.l.Warn("两次密码不一致", logger.String("email", req.Email))
		return
	}
	if errors.Is(err, user_domain.ErrTheNicknameIsTooLong) {
		ctx.JSON(http.StatusOK, web.Result{
			Code: 1,
			Msg:  "昵称长度为 3 - 20 个字",
			Data: nil,
		})
		a.l.Warn("昵称长度为 3 - 20 个字", logger.String("email", req.Email))
		return
	}
	if errors.Is(err, user_domain.ErrTheMobilePhoneNumberIsInvalid) {
		ctx.JSON(http.StatusOK, web.Result{
			Code: 1,
			Msg:  "手机号码无效",
			Data: nil,
		})
		a.l.Warn("手机号码无效", logger.String("phone", req.Phone))
		return
	}
	if errors.Is(err, user_service.ErrEmailAlreadyExists) {
		ctx.JSON(http.StatusOK, web.Result{
			Code: 1,
			Msg:  "邮箱已存在",
			Data: nil,
		})
		a.l.Warn("邮箱已存在", logger.String("email", req.Email))
		return
	}
	if errors.Is(err, user_service.ErrPhoneAlreadyExists) {
		ctx.JSON(http.StatusOK, web.Result{
			Code: 1,
			Msg:  "手机号已存在",
			Data: nil,
		})
		a.l.Warn("手机号已存在", logger.String("phone", req.Phone))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, web.Result{
			Code: 2,
			Msg:  "系统错误",
			Data: nil,
		})
		a.l.Error("系统错误", logger.Error("err", err))
		return
	}

	ctx.JSON(http.StatusOK, web.Result{
		Code: 0,
		Msg:  "管理员注册成功",
		Data: nil,
	})
	a.l.Info("管理员注册成功", logger.String("email", req.Email))
}
