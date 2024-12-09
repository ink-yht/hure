package user_service

import (
	"context"
	"errors"
	"github.com/ink-yht/hure/internal/domain/user_domain"
	"github.com/ink-yht/hure/internal/repository/user_repo"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailAlreadyExists = errors.New("电子邮件已存在")
	ErrPhoneAlreadyExists = errors.New("电话已存在")
)

// AdminService 定义了用户服务的接口
type AdminService interface {
	Signup(ctx context.Context, req user_domain.AdminRegisterRequest) error
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

func (svc *AdminServiceImpl) Signup(ctx context.Context, req user_domain.AdminRegisterRequest) error {
	// 校验请求
	if err := req.Validate(); err != nil {
		return err
	}

	// 密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		// 密码加密失败
		return err
	}

	err = svc.repo.Create(ctx, user_domain.Admin{
		Email:     req.Email,
		Password:  string(hash),
		Phone:     req.Phone,
		Nickname:  req.Nickname,
		Avatar:    req.Avatar,
		Signature: req.Signature,
	})
	if err != nil {
		if errors.Is(err, user_repo.ErrDuplicateEmail) {
			return ErrEmailAlreadyExists
		}
		if errors.Is(err, user_repo.ErrDuplicatePhone) {
			return ErrPhoneAlreadyExists
		}
		return err
	}
	return nil
}
