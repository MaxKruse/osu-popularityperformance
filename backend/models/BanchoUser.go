package models

type BanchoUser struct {
	AvatarURL string `json:"avatar_url"`
	Id        int    `json:"id"`
	Username  string `json:"username"`
}
