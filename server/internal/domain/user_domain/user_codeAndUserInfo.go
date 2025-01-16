package user_domain

// CodeAndUserInfoRequest 登录凭证code和用户信息请求体
type CodeAndUserInfoRequest struct {
	Code      string `json:"code"`
	AvatarUrl string `json:"avatarUrl"`
	NickName  string `json:"nickName"`
	Gender    int8   `json:"gender"`
}

// RoleRequest 登录凭证code和用户信息请求体
type RoleRequest struct {
	ID   uint `json:"id"`
	Role int8 `json:"role"`
}
