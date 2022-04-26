package mysql

import (
	"context"
	"fmt"

	"pluseid.io/invitation/data/entity/model"
)

func (m *mysql) GetAllInvitations(ctx context.Context) ([]model.Invitation, error) {
	lists := []model.Invitation{}
	rows, err := m.db.Query("SELECT id, code, expire FROM invitations")
	if err != nil {
		return lists, err
	}
	defer rows.Close()
	for rows.Next() {
		var item model.Invitation
		err := rows.Scan(&item.ID, &item.Code, &item.Expire)
		if err != nil {
			fmt.Println(err)
		}
		lists = append(lists, item)
	}
	return lists, nil
}

func (m *mysql) CreateNewInviteCode(ctx context.Context, user *model.User) error {
	// row := db.QueryRow("INSERT * value ()", birdName, 1)
	return nil
}
