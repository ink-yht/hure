package user_dao

import (
	"database/sql"
)

type UserDao interface {
}

type UserDAOImpl struct {
	db *sql.DB
}

func NewUserDAO(db *sql.DB) UserDao {
	return &UserDAOImpl{db: db}
}
