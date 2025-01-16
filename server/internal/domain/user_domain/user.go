package user_domain

type User struct {
	ID        uint   `json:"id"`         // id
	OpenID    string `json:"openid"`     // 微信openid
	Nickname  string `json:"nickname"`   // 微信昵称
	Avatar    string `json:"avatar"`     // 微信头像
	Phone     string `json:"phone"`      // 手机号
	Gender    int8   `json:"gender"`     // 性别（男、女、未知） 1 男，2 女，3 未知
	Role      int8   `json:"role"`       // 当前角色（求职者、招聘者） 1 求职者 2 招聘者
	Status    int8   `son:"status"`      // 用户状态（激活、封禁） 1 激活，2 封禁
	CreatedAt int64  `json:"created_at"` // 创建时间
	UpdatedAt int64  `json:"updated_at"` // 更新时间
}

// GetSexText 响应给前端：在返回数据时，将数值转为文本
func GetSexText(sex int8) string {
	switch sex {
	case 1:
		return "男"
	case 2:
		return "女"
	default:
		return "未知"
	}
}

// GetSexValue 后端接受处理时解析为数值存储
func GetSexValue(sexText string) int8 {
	switch sexText {
	case "男":
		return 1
	case "女":
		return 2
	default:
		return 99 // 未知或未指定
	}
}
