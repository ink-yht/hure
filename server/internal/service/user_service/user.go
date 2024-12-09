package user_service

import (
	"github.com/ink-yht/hure/internal/repository/user_repo"
)

// UserService 定义了用户服务的接口
type UserService interface {
}

// UserServiceImpl 实现了 UserService 接口
type UserServiceImpl struct {
	repo user_repo.UserRepository
}

func NewUserService(repo user_repo.UserRepository) UserService {
	return &UserServiceImpl{
		repo: repo,
	}
}
