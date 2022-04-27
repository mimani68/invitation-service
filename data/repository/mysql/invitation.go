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

func (m *mysql) GetInvitationById(ctx context.Context, invitationId int) (model.Invitation, error) {
	item := model.Invitation{}
	rows, err := m.db.Query("SELECT id, code, expire FROM invitations WHERE id = ?", invitationId)
	if err != nil {
		return item, err
	}
	defer rows.Close()
	errScan := rows.Scan(&item.ID, &item.Code, &item.Expire)
	if err != nil {
		fmt.Println(errScan)
	}
	if err != nil {
		return item, err
	}
	return item, nil
}

func (m *mysql) CreateNewInviteCode(ctx context.Context, newItem model.Invitation) (model.Invitation, error) {
	sqlStatement := `
	INSERT INTO invitations (code, expire, is_active, include, exclude)
	VALUES ($1, $2, true, $3, $4)
	RETURNING id;
	`
	result, err := m.db.Exec(sqlStatement, newItem.Code, newItem.Expire, newItem.Include, newItem.Exclude)
	if err != nil {
		return newItem, err
	}
	fmt.Println(result)
	// errScan := rows.Scan(&item.ID, &item.Code, &item.Expire)
	// if err != nil {
	// 	fmt.Println(errScan)
	// }
	if err != nil {
		return newItem, err
	}
	return newItem, nil
}

func (m *mysql) UpdateInvitation(ctx context.Context, invitationId int, updatedItem model.Invitation) (model.Invitation, error) {
	sqlStatement := `
	UPDATE invitations
	SET code = $2, expire = $3, include = $4, exclude = $5
	WHERE id = $1
	RETURNING id, code;
	`
	result, err := m.db.Exec(sqlStatement, invitationId, updatedItem.Code, updatedItem.Expire, updatedItem.Include, updatedItem.Exclude)
	if err != nil {
		return updatedItem, err
	}
	fmt.Println(result)
	// errScan := rows.Scan(&item.ID, &item.Code, &item.Expire)
	// if err != nil {
	// 	fmt.Println(errScan)
	// }
	if err != nil {
		return updatedItem, err
	}
	return updatedItem, nil
}

func (m *mysql) DeleteInvitation(ctx context.Context, invitationId int) (model.Invitation, error) {
	sqlStatement := `
	DELETE FROM invitations
	WHERE id = $1
	RETURNING id, code;
	`
	result, err := m.db.Exec(sqlStatement, invitationId)
	if err != nil {
		return model.Invitation{}, err
	}
	fmt.Println(result)
	// errScan := rows.Scan(&item.ID, &item.Code, &item.Expire)
	// if err != nil {
	// 	fmt.Println(errScan)
	// }
	if err != nil {
		return model.Invitation{}, err
	}
	return model.Invitation{}, nil
}
