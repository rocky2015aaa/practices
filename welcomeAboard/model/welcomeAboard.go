package model

import (
	"time"
)

type CookieData struct {
	UserName    string    `json:"user_name"`
	EncryptedId string    `json:"encrypted_id"`
	Expiration  time.Time `json:"expiration"`
}

type CookieForCache struct {
	EncryptedId string    `json:"encrypted_id"`
	Expiration  time.Time `json:"expiration"`
}

type LoginData struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
