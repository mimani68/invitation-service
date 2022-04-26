package model

import (
	"time"
)

type Invitation struct {
	ID               int       `json:"id"`
	Code             string    `json:"code"`
	Expire           string    `json:"expire"`
	Include          string    `json:"include"`
	Exclude          string    `json:"exclude"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at"`
	DeletedDescption string    `json:"delete_desc"`
}
