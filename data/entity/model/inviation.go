package model

type Invitation struct {
	ID               int    `json:"id"`
	Code             string `json:"code"`
	Expire           string `json:"expire"`
	Include          string `json:"include"`
	Exclude          string `json:"exclude"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	DeletedDescption string `json:"delete_desc"`
}
