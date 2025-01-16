package user_dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type UserDao interface {
	WithTransaction(ctx context.Context, f func(tx *sql.Tx) error) error
	Insert(ctx context.Context, tx *sql.Tx, user User) error
	FindByOpenID(ctx context.Context, openid string) (User, error)
	UpdateInfo(ctx context.Context, user User) error
}

type SqlUserDAO struct {
	db *sql.DB
}

func NewUserDAO(db *sql.DB) UserDao {
	return &SqlUserDAO{db: db}
}

func (dao *SqlUserDAO) UpdateInfo(ctx context.Context, user User) error {

	// 构建更新语句，这里假设admin结构体中有对应的ID字段用于定位要更新的记录
	query := `
        UPDATE user
        SET role = ?, updated_at = ?
        WHERE id =?
    `

	// 准备参数
	args := []interface{}{
		user.Role,
		user.UpdatedAt,
		user.ID,
	}

	// 执行更新操作
	result, err := dao.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	// 检查受影响的行数（可选）
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")
	}

	return nil
}

func (dao *SqlUserDAO) FindByOpenID(ctx context.Context, openid string) (User, error) {

	var user User

	// SQL查询语句，用于从admins表中查询指定邮箱的记录
	query := `
        SELECT id, openid, nickname, avatar, phone, gender, role, status, created_at, updated_at
        FROM user
        WHERE openid = ?
    `

	// 执行SQL查询并解析结果到Admin结构体中
	err := dao.db.QueryRowContext(ctx, query, openid).Scan(
		&user.ID,
		&user.OpenID, &user.Nickname, &user.Avatar,
		&user.Phone, &user.Gender, &user.Role, &user.Status,
		&user.CreatedAt, &user.UpdatedAt,
	)

	// 如果解析过程中发生错误，进一步处理错误
	if err != nil {
		// 如果错误类型为sql.ErrNoRows，表示没有找到任何记录
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrTheUserDoesNotExist // 没有找到记录，返回空的Admin结构体
		}
		// 如果是其他类型的错误，直接返回错误信息
		return User{}, err
	}

	// 查询成功，返回解析出的Admin结构体
	return user, nil
}

func (dao *SqlUserDAO) Insert(ctx context.Context, tx *sql.Tx, user User) error {
	query := `
        INSERT INTO user (openid, nickname, avatar, phone, gender, role, status, created_at, updated_at)
        VALUES (?,?,?,?,?,?,?,?,?)
    `
	_, err := tx.ExecContext(ctx, query,
		user.OpenID, user.Nickname, user.Avatar,
		user.Phone, user.Gender, user.Role, user.Status,
		user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		// 直接返回原始错误
		return err
	}
	// 插入成功，返回nil
	return nil
}

// WithTransaction 包裹事务逻辑
// 该方法用于执行需要事务处理的操作，接受一个操作函数作为参数
// 如果事务开始失败，返回错误信息
// 如果操作函数执行成功且没有发生错误，提交事务；否则回滚事务
func (dao *SqlUserDAO) WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
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
