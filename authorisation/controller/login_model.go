package controller

import (
	"time"
)

type User struct {
	Id int `json:"id"`
	*LoginInfo
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
