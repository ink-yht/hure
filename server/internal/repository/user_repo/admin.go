package user_repo

import (
	"context"
	"database/sql"
	"github.com/ink-yht/hure/internal/domain/user_domain"
	"github.com/ink-yht/hure/internal/repository/dao/user_dao"
)

var (
	ErrDuplicateEmail      = user_dao.ErrDuplicateEmail
	ErrDuplicatePhone      = user_dao.ErrDuplicatePhone
	ErrTheUserDoesNotExist = user_dao.ErrTheUserDoesNotExist
)

type AdminRepository interface {
	Create(ctx context.Context, admin user_domain.Admin) error
	FindByEmail(ctx context.Context, email string) (user_domain.Admin, error)
	FindById(ctx context.Context, id uint) (user_domain.Admin, error)
	UpdateInfo(ctx context.Context, admin user_domain.Admin) error
}

type AdminRepositoryImpl struct {
	dao user_dao.AdminDao
}

func NewAdminRepository(dao user_dao.AdminDao) AdminRepository {
	return &AdminRepositoryImpl{
		dao: dao,
	}
}

// UpdateInfo 更新管理员信息
// 该方法接收一个上下文和一个管理员对象作为参数
// 使用事务来确保更新操作的完整性
func (repo *AdminRepositoryImpl) UpdateInfo(ctx context.Context, admin user_domain.Admin) error {
	// 包裹事务
	return repo.dao.WithTransaction(ctx, func(tx *sql.Tx) error {
		// 插入数据库
		err := repo.dao.UpdateInfo(ctx, tx, repo.domainToEntity(admin))
		if err != nil {
			return err
		}

		// 缓存管理员
		//return r.AdminCache.SetAdmin(ctx, admin.Email, daoAdmin.ID)
		return nil
	})
}

// FindById 根据用户ID查找管理员信息。
// 该方法使用上下文对象ctx来取消请求、设置请求的截止时间、传递请求范围的值等。
// 参数:
//
//	ctx - 上下文对象，用于传递请求范围的值、取消请求等。
//	id - 管理员的唯一标识符。
//
// 返回值:
//
//	user_domain.Admin - 查找到的管理员领域对象。
//	error - 如果查找过程中发生错误，返回错误信息。
func (repo *AdminRepositoryImpl) FindById(ctx context.Context, id uint) (user_domain.Admin, error) {
	// 调用数据访问对象dao的FindById方法查找数据库中的管理员信息。
	daoUser, err := repo.dao.FindById(ctx, id)
	// 如果查找过程中发生错误，返回一个空的管理员领域对象和错误信息。
	if err != nil {
		return user_domain.Admin{}, err
	}
	// 将查找到的管理员数据访问对象转换为领域对象，并返回。
	return repo.entityToDomain(daoUser), nil
}

// FindByEmail 通过电子邮件地址查找管理员信息。
// 该方法接收一个上下文和一个电子邮件字符串作为参数，
// 并返回一个管理员实体或错误。
// 主要用途是通过电子邮件地址检索管理员的详细信息。
func (repo *AdminRepositoryImpl) FindByEmail(ctx context.Context, email string) (user_domain.Admin, error) {
	// 调用DAO层的FindByEmail方法来查找数据库中的用户。
	daoUser, err := repo.dao.FindByEmail(ctx, email)
	// 如果发生错误，返回一个空的管理员实体和错误。
	if err != nil {
		return user_domain.Admin{}, err
	}
	// 将DAO层返回的用户实体转换为领域模型，并返回。
	return repo.entityToDomain(daoUser), nil
}

// Create 在数据库中创建一个新的管理员账户
// 该方法使用事务来确保数据的一致性
// 参数:
//
//	ctx - 上下文，用于传递请求范围的数据、取消信号等
//	admin - 要创建的管理员对象
//
// 返回值:
//
//	error - 如果创建过程中发生错误，则返回该错误
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

// domainToEntity 将管理员领域模型转换为数据访问层实体
// 参数:
//
//	u - 管理员领域模型
//
// 返回值:
//
//	user_dao.Admin - 转换后的数据访问层实体
func (repo *AdminRepositoryImpl) domainToEntity(u user_domain.Admin) user_dao.Admin {
	return user_dao.Admin{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
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

// entityToDomain 将数据访问层实体转换为管理员领域模型
// 参数:
//
//	u - 数据访问层实体
//
// 返回值:
//
//	user_domain.Admin - 转换后的管理员领域模型
func (repo *AdminRepositoryImpl) entityToDomain(u user_dao.Admin) user_domain.Admin {
	return user_domain.Admin{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Email:     nullStringToString(u.Email),
		Phone:     nullStringToString(u.Phone),
		Password:  u.Password,
		Nickname:  u.Nickname,
		Signature: u.Signature,
		Avatar:    u.Avatar,
	}
}

// 辅助函数：将 sql.NullString 转换为 string
// 参数:
//
//	ns - 可能为空的字符串
//
// 返回值:
//
//	string - 如果 ns.Valid 为 true，则返回 ns.String，否则返回空字符串
func nullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
