package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"pluseid.io/invitation/data/entity/model"
	"pluseid.io/invitation/data/repository"
)

func (m *mysql) GetAllInvitations(ctx context.Context, meta repository.QueryMeta) ([]model.Invitation, error) {
	lists := []model.Invitation{}
	var rows *sql.Rows
	var err error
	if meta.Query {
		sql := "SELECT id, code, expire FROM invitations WHERE is_active = ?"
		rows, err = m.db.Query(sql, meta.IsActive)
	} else {
		sql := "SELECT id, code, expire FROM invitations"
		rows, err = m.db.Query(sql)
	}
	if err != nil {
		return lists, err
	}
	defer rows.Close()
	for rows.Next() {
		var item model.Invitation
		err := rows.Scan(&item.Id, &item.Code, &item.Expire)
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
	rows.Next()
	errScan := rows.Scan(&item.Id, &item.Code, &item.Expire)
	if err != nil {
		fmt.Println(errScan)
	}
	if err != nil {
		return item, err
	}
	return item, nil
}

func (m *mysql) GetInvitationByCode(ctx context.Context, code string) (bool, error) {
	item := model.Invitation{}
	sqlStatement := `SELECT id, code, expire FROM invitations WHERE code = ? AND expire >= NOW()`
	rows, err := m.db.Query(sqlStatement, code)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	rows.Next()
	errScan := rows.Scan(&item.Id, &item.Code, &item.Expire)
	if errScan != nil {
		fmt.Println(errScan)
		return false, err
	}
	return item.Id > 0, nil
}

func (m *mysql) CreateNewInviteCode(ctx context.Context, newItem model.Invitation) (model.Invitation, error) {
	sqlStatement := `
	INSERT INTO invitations (id, code, expire, is_active, include, exclude, created_at)
	VALUES (NULL, ?, ?, true, ?, ?, NOW());
	`
	item := model.Invitation{}
	rows, err := m.db.Query(sqlStatement, newItem.Code, newItem.Expire, newItem.Include, newItem.Exclude)
	if err != nil {
		return item, err
	}
	defer rows.Close()
	rows.Next()
	errScan := rows.Scan(&item.Id, &item.CreatedAt)
	if errScan != nil {
		fmt.Println(errScan)
	}
	newItem.Id = item.Id
	newItem.CreatedAt = item.CreatedAt
	return newItem, nil
}

func (m *mysql) UpdateInvitation(ctx context.Context, invitationId int, updatedItem model.Invitation) (model.Invitation, error) {
	sqlStatement := `
	UPDATE invitations
	SET code = ?, expire = ?, include = ?, exclude = ?, is_active = ?
	WHERE id = ?;
	`
	item := model.Invitation{}
	rows, err := m.db.Query(sqlStatement, updatedItem.Code, updatedItem.Expire, updatedItem.Include, updatedItem.Exclude, updatedItem.IsActive, invitationId)
	if err != nil {
		return updatedItem, err
	}
	defer rows.Close()
	rows.Next()
	errScan := rows.Scan(&item.Id, &item.CreatedAt)
	if errScan != nil {
		fmt.Println(errScan)
	}
	return updatedItem, nil
}

func (m *mysql) DeleteInvitation(ctx context.Context, invitationId int) (model.Invitation, error) {
	sqlStatement := `
	DELETE FROM invitations
	WHERE id = ?;
	`
	item := model.Invitation{}
	rows, err := m.db.Query(sqlStatement, invitationId)
	if err != nil {
		return model.Invitation{}, err
	}
	defer rows.Close()
	rows.Next()
	errScan := rows.Scan(&item.Id, &item.CreatedAt)
	if errScan != nil {
		fmt.Println(errScan)
	}
	return model.Invitation{}, nil
}
