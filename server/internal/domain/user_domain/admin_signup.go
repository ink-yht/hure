package user_domain

import (
	"errors"
	"github.com/dlclark/regexp2"
)

var (
	emailRegex                          = regexp2.MustCompile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, regexp2.None)
	passwordRegex                       = regexp2.MustCompile(`^(?=.*[a-zA-Z])(?=.*[0-9])(?=.*[._~!@#$^&*])[A-Za-z0-9._~!@#$^&*]{8,20}$`, regexp2.None)
	phoneRegex                          = regexp2.MustCompile(`^1[3-9]\d{9}$`, regexp2.None)
	nicknameRegex                       = regexp2.MustCompile(`^[\u4e00 - \u9fffA - Za - z0 - 9\-_]{3,20}$`, regexp2.None)
	signatureRegex                      = regexp2.MustCompile(`^.{0,100}$`, regexp2.None)
	ErrTheMailboxIsNotInTheRightFormat  = errors.New("电子邮件格式无效")
	ErrThePasswordIsNotInTheRightFormat = errors.New("密码长度必须为 8-20 个字符，并包含字母、数字和特殊字符")
	ErrThePasswordIsInconsistentTwice   = errors.New("两次密码不一致")
	ErrTheMobilePhoneNumberIsInvalid    = errors.New("手机号码无效")
	ErrTheNicknameIsTooLong             = errors.New("昵称长度为 3 - 20 个字")
)

// AdminRegisterRequest 用户邮箱注册请求体
type AdminRegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Phone           string `json:"phone"`
	Avatar          string `json:"avatar"`
	Nickname        string `json:"nickname"`
	Signature       string `json:"signature"`
}

// Validate 校验请求参数
func (req *AdminRegisterRequest) Validate() error {
	// 校验邮箱格式
	if match, _ := emailRegex.MatchString(req.Email); !match {
		return ErrTheMailboxIsNotInTheRightFormat
	}

	// 校验密码格式
	if match, _ := passwordRegex.MatchString(req.Password); !match {
		return ErrThePasswordIsNotInTheRightFormat
	}

	// 确认密码是否一致
	if req.Password != req.ConfirmPassword {
		return ErrThePasswordIsInconsistentTwice
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
