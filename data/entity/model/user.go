package model

import (
	"time"

	"pluseid.io/invitation/data/entity/enum"
)

type User struct {
	ID        int       `json:"id"`
	Username  int       `json:"username"`
	Password  string    `json:"-"`
	Role      enum.Role `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
