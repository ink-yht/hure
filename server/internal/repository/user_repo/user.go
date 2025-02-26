package user_repo

import (
	"github.com/ink-yht/hure/internal/repository/dao/user_dao"
)

type UserRepository interface {
}

type UserRepositoryImpl struct {
	dao user_dao.UserDao
}

func NewUserRepository(dao user_dao.UserDao) UserRepository {
	return &UserRepositoryImpl{
		dao: dao,
	}
}
