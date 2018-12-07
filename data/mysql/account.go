package mysql

import (
	"boyter/portfold/data"
	"database/sql"
)

type AccountModel struct {
	DB *sql.DB
}

func (m *AccountModel) Insert(project data.Project) (*data.Account, error) {
	return nil, nil
}

func (m *AccountModel) GetAccount(id int) (*data.Account, error) {
	stmt := `
		SELECT	id,
             	name,
             	email,
				created,
				updated,
				active,
				details
		FROM	account
		WHERE	id = ?
		  AND   active = 1
`
	row := m.DB.QueryRow(stmt, id)
	account := &data.Account{}

	err := row.Scan(&account.Id, &account.Name, &account.Created, &account.Updated)
	if err == sql.ErrNoRows {
		return nil, data.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return account, nil
}
