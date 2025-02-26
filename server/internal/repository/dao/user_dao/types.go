package user_dao

import "database/sql"

// User 用户
type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`    // id
	OpenID    string `gorm:"size:50;unique;not null" json:"openid"` // 微信openid
	Nickname  string `gorm:"size:100" json:"nickname"`              // 微信昵称
	Avatar    string `gorm:"size:255" json:"avatar"`                // 微信头像
	Phone     string `gorm:"size:11" json:"phone"`                  // 手机号
	Gender    int8   `gorm:"size:10" json:"gender"`                 // 性别（男、女、未知） 1 男，2 女，3 未知
	Role      int8   `gorm:"size:10;default:'1'" json:"role"`       // 当前角色（求职者、招聘者） 1 求职者 2 招聘者
	Status    int8   `gorm:"size:10;default:'1'" json:"status"`     // 用户状态（激活、封禁） 1 激活，2 封禁
	CreatedAt int64  `json:"created_at"`                            // 创建时间
	UpdatedAt int64  `json:"updated_at"`                            // 更新时间
}

// Resume 简历表
type Resume struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"` // ID
	UserID    uint   `gorm:"not null;unique" json:"user_id"`     // 用户 ID 外键
	Content   string `gorm:"type:text" json:"content"`           // 简历内容（文本或 JSON 格式）
	FilePath  string `gorm:"size:255" json:"file_path"`          // 简历文件路径（上传的简历文件）
	CreatedAt int64  `json:"created_at"`                         // 创建时间
	UpdatedAt int64  `json:"updated_at"`                         // 更新时间
}

// BusinessLicense 招聘者营业执照表
type BusinessLicense struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"` // ID
	UserID    uint   `gorm:"not null;unique" json:"user_id"`     // 用户 ID 外键
	Content   string `gorm:"type:text" json:"content"`           // 营业执照信息
	FilePath  string `gorm:"size:255" json:"file_path"`          // 营业执照文件路径
	CreatedAt int64  `json:"created_at"`                         // 创建时间
	UpdatedAt int64  `json:"updated_at"`                         // 更新时间
}

// RealNameVerification 实名审核表
type RealNameVerification struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"` // ID
	UserID         uint   `gorm:"not null;unique" json:"user_id"`     // 用户 ID 外键
	RealName       string `gorm:"size:50" json:"real_name"`           // 真实姓名
	IDCardNumber   string `gorm:"size:18" json:"id_card_number"`      // 身份证号码
	IDCardFrontImg string `gorm:"size:255" json:"id_card_front_img"`  // 身份证正面图片路径
	IDCardBackImg  string `gorm:"size:255" json:"id_card_back_img"`   // 身份证背面图片路径
	VerifiedStatus int8   `gorm:"default:0" json:"verified_status"`   // 审核状态（0 待审核，1 通过，2 拒绝）
	VerifiedTime   int64  `json:"verified_time"`                      // 审核时间
	CreatedAt      int64  `json:"created_at"`                         // 创建时间
	UpdatedAt      int64  `json:"updated_at"`                         // 更新时间
}

// Admin 管理员
type Admin struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"` // id
	Email     sql.NullString `gorm:"size:100;unique" json:"email"`       // 邮箱
	Password  string         `gorm:"not null" json:"password"`           // 加密后的密码
	Phone     sql.NullString `gorm:"size:20;unique" json:"phone"`        // 手机号
	Avatar    string         `gorm:"size:255" json:"avatar"`             // 头像
	Nickname  string         `gorm:"size:32"`                            // 昵称
	Signature string         `gorm:"size:128"`                           // 个性签名
	CreatedAt int64          `json:"created_at"`                         // 创建时间
	UpdatedAt int64          `json:"updated_at"`                         // 更新时间
}
