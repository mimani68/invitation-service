package mysql

import (
	"context"
	"fmt"

	"pluseid.io/invitation/data/entity/model"
)

func (m *mysql) GetAllInvitations(ctx context.Context) ([]model.Invitation, error) {
	// rows := m.db.QueryRow("SELECT * FROM invitation WHERE bird = $1 LIMIT $2", birdName, 1)
	lists := []model.Invitation{}
	err := m.db.QueryRow("SELECT * FROM invitations LIMIT $1", 10).Scan(&lists)
	if err != nil {
		fmt.Println(err)
	}
	return lists, nil
}

func (m *mysql) CreateNewInviteCode(ctx context.Context, user *model.User) error {
	// row := db.QueryRow("INSERT * value ()", birdName, 1)
	return nil
}
