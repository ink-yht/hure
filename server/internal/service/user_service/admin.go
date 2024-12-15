package user_service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ink-yht/hure/internal/domain/user_domain"
	"github.com/ink-yht/hure/internal/repository/user_repo"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	ErrEmailAlreadyExists  = errors.New("电子邮件已存在")
	ErrPhoneAlreadyExists  = errors.New("电话已存在")
	ErrTheUserDoesNotExist = errors.New("用户不存在")
)

// AdminService 定义了用户服务的接口
type AdminService interface {
	Signup(ctx context.Context, req user_domain.AdminRegisterRequest) error
	Login(ctx context.Context, req user_domain.AdminLoginRequest, userAgent string) (string, error)
	Info(ctx context.Context, id uint) (user_domain.Admin, error)
	Edit(ctx context.Context, req user_domain.AdminEditRequest) error
}

// AdminServiceImpl 实现了 AdminService 接口
type AdminServiceImpl struct {
	repo user_repo.AdminRepository
}

func NewAdminService(repo user_repo.AdminRepository) AdminService {
	return &AdminServiceImpl{
		repo: repo,
	}
}

// Edit 是 AdminServiceImpl 结构体的一个方法，用于编辑管理员信息。
// 它接收一个 context.Context 类型的上下文参数和一个 user_domain.AdminEditRequest 类型的请求对象。
// 方法通过更新管理员的电子邮件、昵称、电话、签名和最后更新时间来编辑管理员信息。
// 返回值是一个错误类型，用于指示操作是否成功。
func (svc *AdminServiceImpl) Edit(ctx context.Context, req user_domain.AdminEditRequest) error {
	// 校验登录请求的合法性，如检查邮箱格式和密码强度。
	if err := req.EditValidate(); err != nil {
		return err
	}

	// 获取当前时间戳（毫秒）
	now := time.Now().UnixMilli()

	// 调用 repo 的 UpdateInfo 方法更新管理员信息。
	// 这里将 req 中的管理员信息与当前时间戳一起封装成一个 Admin 对象传递给 UpdateInfo 方法。
	return svc.repo.UpdateInfo(ctx, user_domain.Admin{
		ID:        req.ID,
		Email:     req.Email,
		Nickname:  req.Nickname,
		Phone:     req.Phone,
		Signature: req.Signature,
		UpdatedAt: now,
	})
}

// Info 获取指定ID的管理员信息
// 该方法通过管理员ID查询数据库，返回相应的管理员对象和错误信息（如果有）。
// 主要用途是为管理员信息的查询提供一个简单、高效的接口。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的上下文信息，如请求ID、用户信息等。
//	id - 管理员的唯一标识符，用于指定需要查询的管理员。
//
// 返回值:
//
//	user_domain.Admin - 查询到的管理员对象，如果未找到，则返回空对象。
//	error - 如果查询过程中发生错误，则返回错误信息。
func (svc *AdminServiceImpl) Info(ctx context.Context, id uint) (user_domain.Admin, error) {
	return svc.repo.FindById(ctx, id)
}

// Login 实现了 AdminService 接口的 Login 方法。
// 它负责处理管理员的登录请求，验证用户凭据并生成访问令牌。
// 参数:
//   - ctx: 上下文，用于传递请求范围的 deadline、取消信号、请求级值。
//   - req: 包含登录所需信息的请求对象，包括邮箱和密码。
//   - userAgent: 客户端的用户代理字符串，用于在生成 JWT 时可能的定制化。
//
// 返回值:
//   - string: 成功时返回访问令牌。
//   - error: 遇到错误时返回错误信息，如验证失败或用户不存在。
func (svc *AdminServiceImpl) Login(ctx context.Context, req user_domain.AdminLoginRequest, userAgent string) (string, error) {
	// 校验登录请求的合法性，如检查邮箱格式和密码强度。
	if err := req.LoginValidate(); err != nil {
		return "", err
	}

	// 从数据库中查找用户，根据提供的邮箱地址。
	user, err := svc.repo.FindByEmail(ctx, req.Email)
	// 处理用户不存在的情况。
	if errors.Is(err, user_repo.ErrTheUserDoesNotExist) {
		return "", ErrTheUserDoesNotExist
	}
	// 处理其他查找错误。
	if err != nil {
		return "", err
	}

	// 校验密码，使用 bcrypt 来比较数据库中的密码哈希和用户提供的密码。
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", ErrTheUserDoesNotExist
	}

	// 生成 JWT，根据用户 ID 和用户代理字符串定制化生成访问令牌。
	token, err := svc.setJWTToken(user.ID, userAgent)
	if err != nil {
		return "", err
	}
	// 返回生成的访问令牌。
	return token, nil
}

// Signup 管理员注册函数
// 该函数接收一个 AdminRegisterRequest 类型的注册请求，并对请求进行处理
// 主要功能包括：请求校验、密码加密和用户信息创建
func (svc *AdminServiceImpl) Signup(ctx context.Context, req user_domain.AdminRegisterRequest) error {
	// 校验请求
	// 对注册请求进行合法性校验，确保数据的完整性和正确性
	if err := req.RegisterValidate(); err != nil {
		return err
	}

	// 密码加密
	// 使用 bcrypt 算法对用户密码进行加密，提高数据安全性
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		// 密码加密失败
		return err
	}

	now := time.Now().UnixMilli()

	// 创建管理员用户
	// 将加密后的密码和其他用户信息一起存入数据库
	err = svc.repo.Create(ctx, user_domain.Admin{
		CreatedAt: now,
		UpdatedAt: now,
		Email:     req.Email,
		Password:  string(hash),
		Phone:     req.Phone,
		Nickname:  req.Nickname,
		Signature: req.Signature,
	})
	if err != nil {
		// 处理可能的错误
		// 如果错误是由于邮箱或电话号码重复引起的，则返回相应的错误
		if errors.Is(err, user_repo.ErrDuplicateEmail) {
			return ErrEmailAlreadyExists
		}
		if errors.Is(err, user_repo.ErrDuplicatePhone) {
			return ErrPhoneAlreadyExists
		}
		// 其他错误
		return err
	}
	// 注册成功，返回 nil 表示操作成功
	return nil
}

// setJWTToken 生成 token
func (svc *AdminServiceImpl) setJWTToken(uid uint, userAgent string) (string, error) {
	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		Id:        uid,
		UserAgent: userAgent,
		RegisteredClaims: jwt.RegisteredClaims{
			// 过期时间设置
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 3)),
		},
	})
	token, err := tokenStr.SignedString(JWTKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
