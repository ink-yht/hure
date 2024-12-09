package user_dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

var (
	ErrDuplicateEmail = errors.New("邮箱冲突")
	ErrDuplicatePhone = errors.New("手机号冲突")
)

type AdminDao interface {
	WithTransaction(ctx context.Context, f func(tx *sql.Tx) error) error
	Insert(ctx context.Context, tx *sql.Tx, admin Admin) error
}

type SqlAdminDAO struct {
	db *sql.DB
}

func NewAdminDAO(db *sql.DB) AdminDao {
	return &SqlAdminDAO{db: db}
}

// Insert 插入管理员数据
func (dao *SqlAdminDAO) Insert(ctx context.Context, tx *sql.Tx, admin Admin) error {
	query := `
		INSERT INTO admins (email, password, phone, avatar, nickname, signature, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := tx.ExecContext(ctx, query,
		admin.Email, admin.Password, admin.Phone,
		admin.Avatar, admin.Nickname, admin.Signature,
		admin.CreatedAt, admin.UpdatedAt,
	)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			if strings.Contains(mysqlErr.Message, "email") {
				return ErrDuplicateEmail
			}
			if strings.Contains(mysqlErr.Message, "phone") {
				return ErrDuplicatePhone
			}
		}
		return err
	}
	return nil
}

// WithTransaction 包裹事务逻辑
func (dao *SqlAdminDAO) WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := dao.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
