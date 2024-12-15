package user_domain

// AdminLoginRequest 用户邮箱登录请求体
type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginValidate 校验请求参数
func (req *AdminLoginRequest) LoginValidate() error {
	// 校验邮箱格式
	if match, _ := emailRegex.MatchString(req.Email); !match {
		return ErrTheMailboxIsNotInTheRightFormat
	}

	// 校验密码格式
	if match, _ := passwordRegex.MatchString(req.Password); !match {
		return ErrThePasswordIsNotInTheRightFormat
	}

	return nil
}
