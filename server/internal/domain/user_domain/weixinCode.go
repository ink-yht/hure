package user_domain

type SessionResponse struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
}
