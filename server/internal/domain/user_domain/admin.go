package user_domain

import "time"

// Admin 管理员
type Admin struct {
	ID        uint      `json:"id"`         // id
	Email     string    `json:"email"`      // 邮箱
	Password  string    `json:"password"`   // 加密后的密码
	Phone     string    `json:"phone"`      // 手机号
	Avatar    string    `json:"avatar"`     // 头像
	Nickname  string    `json:"nickname"`   // 昵称
	Signature string    `json:"signature"`  // 个性签名
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}
