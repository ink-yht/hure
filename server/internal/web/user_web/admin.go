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
	ag.POST("/login", a.Login)   // 管理员登录
	ag.GET("/info", a.Info)      // 管理员个人信息获取
	ag.POST("/edit", a.Edit)     // 管理员个人信息修改
}

// SignUp 管理员注册
// 该方法处理管理员用户的注册请求，验证用户输入信息的正确性并完成注册流程
func (a *AdminHandler) SignUp(ctx *gin.Context) {
	// 解析请求参数
	var req user_domain.AdminRegisterRequest
	err := ctx.Bind(&req)
	if err != nil {
		// 如果参数解析失败，会返回 400 错误
		return
	}

	// 调用服务层方法进行注册，根据返回的错误类型决定响应内容
	err = a.svc.Signup(ctx, req)

	// 根据不同错误类型返回相应错误信息
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

	// 注册成功后返回成功信息
	ctx.JSON(http.StatusOK, web.Result{
		Code: 0,
		Msg:  "管理员注册成功",
		Data: nil,
	})
	a.l.Info("管理员注册成功", logger.String("email", req.Email))
}

// Login 实现管理员的登录逻辑。
// 该方法首先尝试从HTTP请求中解析出管理员登录请求的数据，
// 然后调用服务层进行登录验证，最后根据登录结果返回相应的HTTP响应。
func (a *AdminHandler) Login(ctx *gin.Context) {
	var req user_domain.AdminLoginRequest
	// 将请求体绑定到req变量中，如果解析失败，会自动返回400错误。
	err := ctx.Bind(&req)
	if err != nil {
		// 出错会返回 400 错误
		return
	}

	// 提前从请求头中提取 User-Agent 信息，用于后续可能的日志记录或审计。
	userAgent := ctx.GetHeader("User-Agent")

	// 调用服务层方法进行登录验证，根据返回的错误类型决定响应内容。
	token, err := a.svc.Login(ctx, req, userAgent)
	if errors.Is(err, user_service.ErrTheUserDoesNotExist) {
		// 当用户不存在时，返回自定义的错误信息。
		ctx.JSON(http.StatusOK, web.Result{
			Code: 1,
			Msg:  "邮箱或密码不存在",
			Data: nil,
		})
		// 记录警告日志，包括用户的邮箱信息。
		a.l.Warn("邮箱或密码不存在", logger.String("email", req.Email))
		return
	}
	if err != nil {
		// 当出现其他错误时，返回通用的系统错误信息。
		ctx.JSON(http.StatusOK, web.Result{
			Code: 2,
			Msg:  "系统错误",
			Data: nil,
		})
		// 记录错误日志，包括错误详情。
		a.l.Error("系统错误", logger.Error("err", err))
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
	a.l.Info("登录成功", logger.String("email", req.Email))
}

// Info AdminHandler的Info方法用于获取管理员的个人信息。
// 该方法接收一个gin的Context作为参数，其中包含了请求的相关信息。
// 它从上下文中获取用户声明（claims），并使用这些信息来查询用户详情。
// 如果查询过程中出现错误，它将返回一个表示系统错误的结果。
// 如果查询成功，它将返回一个表示个人信息获取成功的响应，并包含用户信息。
func (a *AdminHandler) Info(ctx *gin.Context) {
	// 从上下文中获取用户声明，转换为UserClaims类型。
	userClaims := ctx.MustGet("claims").(*user_service.UserClaims)

	// 调用服务层的Info方法来获取用户信息。
	user, err := a.svc.Info(ctx, userClaims.Id)
	if err != nil {
		// 如果发生错误，返回一个系统错误的JSON响应，并记录错误日志。
		ctx.JSON(http.StatusOK, web.Result{
			Code: 2,
			Msg:  "系统错误",
			Data: nil,
		})
		a.l.Error("系统错误", logger.Error("err", err))
		return
	}

	// 手动将密码字段设置为空
	user.Password = ""

	// 如果成功获取用户信息，返回一个成功的JSON响应，并记录信息日志。
	ctx.JSON(http.StatusOK, web.Result{
		Code: 0,
		Msg:  "个人信息获取成功",
		Data: user,
	})
	a.l.Info("登录成功", logger.String("email", user.Email))
}

func (a *AdminHandler) Edit(ctx *gin.Context) {
	userClaims := ctx.MustGet("claims").(*user_service.UserClaims)
	var req user_domain.AdminEditRequest
	err := ctx.Bind(&req)
	if err != nil {
		// 出错会返回 400 错误
		return
	}
	req.ID = userClaims.Id
	err = a.svc.Edit(ctx, req)
	// 根据不同错误类型返回相应错误信息
	if errors.Is(err, user_domain.ErrTheMailboxIsNotInTheRightFormat) {
		ctx.JSON(http.StatusOK, web.Result{
			Code: 1,
			Msg:  "电子邮件格式无效",
			Data: nil,
		})
		a.l.Warn("电子邮件格式无效", logger.String("email", req.Email))
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
		Msg:  "个人信息修改成功",
		Data: nil,
	})
	a.l.Info("登录成功", logger.String("email", req.Email))
}
