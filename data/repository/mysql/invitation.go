package mysql

import (
	"context"

	"pluseid.io/invitation/data/entity/model"
)

func (m *mysql) GetAllInvitations(ctx context.Context, user *model.User) error {
	// row := db.QueryRow("SELECT * FROM invitation WHERE bird = $1 LIMIT $2", birdName, 1)
	return nil
}

func (m *mysql) CreateNewInviteCode(ctx context.Context, user *model.User) error {
	// row := db.QueryRow("INSERT * value ()", birdName, 1)
	return nil
}
