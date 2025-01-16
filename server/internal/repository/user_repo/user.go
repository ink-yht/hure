package user_repo

import (
	"context"
	"database/sql"
	"github.com/ink-yht/hure/internal/domain/user_domain"
	"github.com/ink-yht/hure/internal/repository/dao/user_dao"
)

type UserRepository interface {
	Create(ctx context.Context, user user_domain.User) error
	FindByOpenID(ctx context.Context, openid string) (user_domain.User, error)
	UpdateRole(ctx context.Context, user user_domain.User) error
}

type UserRepositoryImpl struct {
	dao user_dao.UserDao
}

func NewUserRepository(dao user_dao.UserDao) UserRepository {
	return &UserRepositoryImpl{
		dao: dao,
	}
}

func (repo *UserRepositoryImpl) UpdateRole(ctx context.Context, user user_domain.User) error {
	return repo.dao.UpdateInfo(ctx, repo.domainToEntity(user))
}

func (repo *UserRepositoryImpl) FindByOpenID(ctx context.Context, openid string) (user_domain.User, error) {
	daoUser, err := repo.dao.FindByOpenID(ctx, openid)
	if err != nil {
		return user_domain.User{}, err
	}
	return repo.entityToDomain(daoUser), nil
}

func (repo *UserRepositoryImpl) Create(ctx context.Context, user user_domain.User) error {
	// 包裹事务
	return repo.dao.WithTransaction(ctx, func(tx *sql.Tx) error {
		// 插入数据库
		err := repo.dao.Insert(ctx, tx, repo.domainToEntity(user))
		if err != nil {
			return err
		}

		// 缓存管理员
		//return r.AdminCache.SetAdmin(ctx, admin.Email, daoAdmin.ID)
		return nil
	})
}

func (repo *UserRepositoryImpl) domainToEntity(u user_domain.User) user_dao.User {
	return user_dao.User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		OpenID:    u.OpenID,
		Phone:     u.Phone,
		Gender:    u.Gender,
		Role:      u.Role,
		Status:    u.Status,
		Nickname:  u.Nickname,
		Avatar:    u.Avatar,
	}
}

func (repo *UserRepositoryImpl) entityToDomain(u user_dao.User) user_domain.User {
	return user_domain.User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		OpenID:    u.OpenID,
		Phone:     u.Phone,
		Gender:    u.Gender,
		Role:      u.Role,
		Status:    u.Status,
		Nickname:  u.Nickname,
		Avatar:    u.Avatar,
	}
}
