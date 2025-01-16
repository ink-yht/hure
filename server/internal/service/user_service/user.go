package user_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ink-yht/hure/internal/domain/user_domain"
	"github.com/ink-yht/hure/internal/repository/user_repo"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	AppID     = "wxda466159775d5757"
	AppSecret = "989da7f2b152c13e2664ac6eb9b4833d"
)

// UserService 定义了用户服务的接口
type UserService interface {
	Login(ctx context.Context, user user_domain.CodeAndUserInfoRequest, userAgent string) (string, error)
	Edit(ctx context.Context, req user_domain.RoleRequest) error
}

// UserServiceImpl 实现了 UserService 接口
type UserServiceImpl struct {
	repo user_repo.UserRepository
}

func NewUserService(repo user_repo.UserRepository) UserService {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (svc *UserServiceImpl) Edit(ctx context.Context, req user_domain.RoleRequest) error {
	now := time.Now().UnixMilli()
	return svc.repo.UpdateRole(ctx, user_domain.User{
		ID:        req.ID,
		UpdatedAt: now,
		Role:      req.Role,
	})
}

func (svc *UserServiceImpl) Login(ctx context.Context, user user_domain.CodeAndUserInfoRequest, userAgent string) (string, error) {
	// code 获取 openID
	sessionInfo, err := svc.GetSessionKeyWithCode(user.Code, AppID, AppSecret)
	fmt.Println(sessionInfo)
	if err != nil {
		return "", err
	}
	fmt.Println("Openid", sessionInfo.Openid)
	// 返回用户信息
	u, err := svc.repo.FindByOpenID(ctx, sessionInfo.Openid)
	if errors.Is(err, user_repo.ErrTheUserDoesNotExist) {
		// 获取当前时间戳（毫秒）
		now := time.Now().UnixMilli()
		// 在数据库中注册
		err = svc.repo.Create(ctx, user_domain.User{
			CreatedAt: now,
			UpdatedAt: now,
			Role:      1,
			Status:    1,
			OpenID:    sessionInfo.Openid,
			Avatar:    user.AvatarUrl,
			Nickname:  user.NickName,
			Gender:    user.Gender,
		})
		if err != nil {
			return "", err
		}

		u, err := svc.repo.FindByOpenID(ctx, sessionInfo.Openid)
		token, err := SetJWTToken(u.ID, userAgent)
		if err != nil {
			return "", err
		}
		// 返回生成的访问令牌。
		return token, nil
	}
	// 处理其他查找错误。
	if err != nil {
		return "", err
	}

	// 用户信息生成 token
	// 生成 JWT，根据用户 ID 和用户代理字符串定制化生成访问令牌。
	token, err := SetJWTToken(u.ID, userAgent)
	if err != nil {
		return "", err
	}
	// 返回生成的访问令牌。
	return token, nil
}

func (svc *UserServiceImpl) GetSessionKeyWithCode(code, appID, appSecret string) (*user_domain.SessionResponse, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session"
	// 构建请求参数
	params := map[string]string{
		"appid":      appID,
		"secret":     appSecret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	query := ""
	for key, value := range params {
		query += key + "=" + value + "&"
	}
	query = query[:len(query)-1] // 去掉末尾多余的&

	fullURL := url + "?" + query

	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	sessionInfo := &user_domain.SessionResponse{}
	err = json.Unmarshal(body, sessionInfo)
	if err != nil {
		return nil, err
	}

	return sessionInfo, nil
}
