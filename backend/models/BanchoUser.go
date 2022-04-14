package models

type BanchoUser struct {
	AvatarURL string `json:"avatar_url"`
	UserId    int    `json:"id"`
	Username  string `json:"username"`
}
