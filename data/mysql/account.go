package mysql

import (
	"boyter/portfold/data"
	"database/sql"
)

type AccountModel struct {
	DB *sql.DB
}

func (m *AccountModel) Delete(account data.Account) error {
	stmt := `DELETE FROM	account 
				   WHERE 	id = ?`
	_, err := m.DB.Exec(stmt, account.Id)
	return err
}

func (m *AccountModel) Insert(account data.Account) (*data.Account, error) {
	stmt := `
		INSERT INTO account(id, name, email, created, updated, active, details)
 		VALUES (NULL,?,?,UTC_TIMESTAMP(),UTC_TIMESTAMP(),1,?)
	`

	res, err := m.DB.Exec(stmt, account.Name, account.Email, account.Details)
	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	// NB potentially an overflow issue here
	acc, err := m.GetAccount(int(lastId))
	if err != nil {
		return nil, err
	}

	return acc, nil
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
`
	row := m.DB.QueryRow(stmt, id)
	account := &data.Account{}

	err := row.Scan(&account.Id, &account.Name, &account.Email, &account.Created, &account.Updated, &account.Active, &account.Details)
	if err == sql.ErrNoRows {
		return nil, data.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return account, nil
}
