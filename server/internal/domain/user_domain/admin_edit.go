package user_domain

// AdminEditRequest 用户信息修改请求体
type AdminEditRequest struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Nickname  string `json:"nickname"`
	Signature string `json:"signature"`
	UpdatedAt int64  `json:"updated_at"` // 更新时间
}

// EditValidate 校验请求参数
func (req *AdminEditRequest) EditValidate() error {
	// 校验邮箱格式
	if match, _ := emailRegex.MatchString(req.Email); !match {
		return ErrTheMailboxIsNotInTheRightFormat
	}

	// 校验手机号格式
	if match, _ := phoneRegex.MatchString(req.Phone); !match {
		return ErrTheMobilePhoneNumberIsInvalid
	}

	// 校验昵称格式
	if match, _ := nicknameRegex.MatchString(req.Nickname); !match {
		return ErrTheNicknameIsTooLong
	}

	// 校验个性签名格式
	if match, _ := signatureRegex.MatchString(req.Signature); !match {
		return ErrTheNicknameIsTooLong
	}

	return nil
}
