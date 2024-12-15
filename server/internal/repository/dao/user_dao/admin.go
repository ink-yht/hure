package user_dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

var (
	ErrDuplicateEmail      = errors.New("邮箱冲突")
	ErrDuplicatePhone      = errors.New("手机号冲突")
	ErrTheUserDoesNotExist = errors.New("用户不存在")
)

type AdminDao interface {
	WithTransaction(ctx context.Context, f func(tx *sql.Tx) error) error
	Insert(ctx context.Context, tx *sql.Tx, admin Admin) error
	FindByEmail(ctx context.Context, email string) (Admin, error)
	FindById(ctx context.Context, id uint) (Admin, error)
	UpdateInfo(ctx context.Context, tx *sql.Tx, admin Admin) error
}

type SqlAdminDAO struct {
	db *sql.DB
}

func NewAdminDAO(db *sql.DB) AdminDao {
	return &SqlAdminDAO{db: db}
}

// UpdateInfo 更新管理员信息
// 该方法使用SQL事务tx来更新数据库中的管理员信息
// 参数:
//
//	ctx - 上下文，用于传递请求范围的值、取消信号等
//	tx - SQL事务，用于执行数据库更新操作
//	admin - 要更新的管理员信息，包括邮箱、电话、昵称、签名等
//
// 返回值:
//
//	如果更新操作成功，返回nil
//	如果更新操作失败，返回相应的错误
func (dao *SqlAdminDAO) UpdateInfo(ctx context.Context, tx *sql.Tx, admin Admin) error {
	// 构建更新语句，这里假设admin结构体中有对应的ID字段用于定位要更新的记录
	query := `
        UPDATE admin
        SET email =?,  phone =?, nickname =?, signature =?, updated_at =?
        WHERE id =?
    `
	// 执行更新操作
	_, err := tx.ExecContext(ctx, query,
		admin.Email, admin.Phone,
		admin.Nickname, admin.Signature,
		admin.UpdatedAt,
		admin.ID, // 使用admin的ID字段来确定要更新的具体记录，这里假设存在ID字段
	)
	if err != nil {
		// 判断错误类型，如果是MySQL的唯一键冲突错误，则进一步处理
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			// 如果错误消息中包含"email"，则认为是邮箱重复错误
			if strings.Contains(mysqlErr.Message, "email") {
				return ErrDuplicateEmail
			}
			// 如果错误消息中包含"phone"，则认为是电话号码重复错误
			if strings.Contains(mysqlErr.Message, "phone") {
				return ErrDuplicatePhone
			}
		}
		// 如果不是唯一键冲突错误，直接返回原始错误
		return err
	}
	// 更新成功，返回nil
	return nil
}

// FindById 根据用户ID查询用户信息。
// 该方法使用SQL查询语句从数据库中获取指定ID的用户信息，并将其映射到Admin结构体中返回。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的上下文信息。
//	id - 用户的唯一标识符。
//
// 返回值:
//
//	Admin - 包含用户信息的结构体。
//	error - 如果查询过程中发生错误，则返回该错误。
func (dao *SqlAdminDAO) FindById(ctx context.Context, id uint) (Admin, error) {
	// 初始化Admin结构体用于存储查询结果
	var admin Admin

	// SQL查询语句，用于从admins表中查询指定邮箱的记录
	query := `
        SELECT id, email, password, phone, avatar, nickname, signature, created_at, updated_at
        FROM admin
        WHERE id = ?
    `

	// 执行SQL查询并解析结果到Admin结构体中
	err := dao.db.QueryRowContext(ctx, query, id).Scan(
		&admin.ID,
		&admin.Email, &admin.Password, &admin.Phone,
		&admin.Avatar, &admin.Nickname, &admin.Signature,
		&admin.CreatedAt, &admin.UpdatedAt,
	)

	// 如果解析过程中发生错误，进一步处理错误
	if err != nil {
		return Admin{}, err
	}

	// 查询成功，返回解析出的Admin结构体
	return admin, nil
}

// FindByEmail 根据邮箱查询管理员数据
// 该方法使用SQL查询来查找指定邮箱的管理员记录
// 如果找到记录，返回对应的Admin结构体和nil；如果没有找到记录或发生错误，返回空的Admin结构体和错误信息
func (dao *SqlAdminDAO) FindByEmail(ctx context.Context, email string) (Admin, error) {
	// 初始化Admin结构体用于存储查询结果
	var admin Admin

	// SQL查询语句，用于从admins表中查询指定邮箱的记录
	query := `
        SELECT id, email, password, phone, avatar, nickname, signature, created_at, updated_at
        FROM admin
        WHERE email = ?
    `

	// 执行SQL查询并解析结果到Admin结构体中
	err := dao.db.QueryRowContext(ctx, query, email).Scan(
		&admin.ID,
		&admin.Email, &admin.Password, &admin.Phone,
		&admin.Avatar, &admin.Nickname, &admin.Signature,
		&admin.CreatedAt, &admin.UpdatedAt,
	)

	// 如果解析过程中发生错误，进一步处理错误
	if err != nil {
		// 如果错误类型为sql.ErrNoRows，表示没有找到任何记录
		if errors.Is(err, sql.ErrNoRows) {
			return Admin{}, ErrTheUserDoesNotExist // 没有找到记录，返回空的Admin结构体
		}
		// 如果是其他类型的错误，直接返回错误信息
		return Admin{}, err
	}

	// 查询成功，返回解析出的Admin结构体
	return admin, nil
}

// Insert 插入管理员数据
// 该方法使用SQL查询来插入一个新的管理员记录
// 如果插入成功，返回nil；如果发生错误，返回错误信息
func (dao *SqlAdminDAO) Insert(ctx context.Context, tx *sql.Tx, admin Admin) error {
	query := `
		INSERT INTO admin (email, password, phone, avatar, nickname, signature, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := tx.ExecContext(ctx, query,
		admin.Email, admin.Password, admin.Phone,
		admin.Avatar, admin.Nickname, admin.Signature,
		admin.CreatedAt, admin.UpdatedAt,
	)
	if err != nil {
		// 判断错误类型，如果是MySQL的唯一键冲突错误，则进一步处理
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			// 如果错误消息中包含"email"，则认为是邮箱重复错误
			if strings.Contains(mysqlErr.Message, "email") {
				return ErrDuplicateEmail
			}
			// 如果错误消息中包含"phone"，则认为是电话号码重复错误
			if strings.Contains(mysqlErr.Message, "phone") {
				return ErrDuplicatePhone
			}
		}
		// 如果不是唯一键冲突错误，直接返回原始错误
		return err
	}
	// 插入成功，返回nil
	return nil
}

// WithTransaction 包裹事务逻辑
// 该方法用于执行需要事务处理的操作，接受一个操作函数作为参数
// 如果事务开始失败，返回错误信息
// 如果操作函数执行成功且没有发生错误，提交事务；否则回滚事务
func (dao *SqlAdminDAO) WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := dao.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		// 如果有panic发生，回滚事务并重新抛出panic
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			// 如果存在错误，回滚事务
			_ = tx.Rollback()
		} else {
			// 否则，提交事务
			err = tx.Commit()
		}
	}()

	// 执行传入的操作函数
	err = fn(tx)
	return err
}
