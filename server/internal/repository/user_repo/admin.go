package user_repo

import (
	"context"
	"database/sql"
	"github.com/ink-yht/hure/internal/domain/user_domain"
	"github.com/ink-yht/hure/internal/repository/dao/user_dao"
	"time"
)

var (
	ErrDuplicateEmail = user_dao.ErrDuplicateEmail
	ErrDuplicatePhone = user_dao.ErrDuplicatePhone
)

type AdminRepository interface {
	Create(ctx context.Context, admin user_domain.Admin) error
}

type AdminRepositoryImpl struct {
	dao user_dao.AdminDao
}

func NewAdminRepository(dao user_dao.AdminDao) AdminRepository {
	return &AdminRepositoryImpl{
		dao: dao,
	}
}

func (repo *AdminRepositoryImpl) Create(ctx context.Context, admin user_domain.Admin) error {
	// 包裹事务
	return repo.dao.WithTransaction(ctx, func(tx *sql.Tx) error {
		// 插入数据库
		err := repo.dao.Insert(ctx, tx, repo.domainToEntity(admin))
		if err != nil {
			return err
		}

		// 缓存管理员
		//return r.AdminCache.SetAdmin(ctx, admin.Email, daoAdmin.ID)
		return nil
	})
}

func (repo *AdminRepositoryImpl) domainToEntity(u user_domain.Admin) user_dao.Admin {
	return user_dao.Admin{
		ID:        u.ID,
		CreatedAt: u.CreatedAt.UnixMilli(),
		UpdatedAt: u.UpdatedAt.UnixMilli(),
		Email: sql.NullString{
			String: u.Email,
			Valid:  len(u.Email) > 0,
		},
		Phone: sql.NullString{
			String: u.Phone,
			Valid:  len(u.Phone) > 0,
		},
		Password:  u.Password,
		Nickname:  u.Nickname,
		Signature: u.Signature,
		Avatar:    u.Avatar,
	}
}

func (repo *AdminRepositoryImpl) entityToDomain(u user_dao.Admin) user_domain.Admin {
	return user_domain.Admin{
		ID:        u.ID,
		CreatedAt: time.UnixMilli(u.CreatedAt),
		UpdatedAt: time.UnixMilli(u.UpdatedAt),
		Email:     nullStringToString(u.Email),
		Phone:     nullStringToString(u.Phone),
		Password:  u.Password,
		Nickname:  u.Nickname,
		Signature: u.Signature,
		Avatar:    u.Avatar,
	}
}

// 辅助函数：将 sql.NullString 转换为 string
func nullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
