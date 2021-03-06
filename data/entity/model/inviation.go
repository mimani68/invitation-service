package model

type Invitation struct {
	Id               int    `json:"id"`
	Code             string `json:"code"`
	Expire           string `json:"expire"`
	Include          string `json:"include"`
	Exclude          string `json:"exclude"`
	IsActive         bool   `json:"isActive" default:"true"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	DeletedDescption string `json:"delete_desc"`
}
