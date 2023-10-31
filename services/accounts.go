package services

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Account struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
	Role     *string `json:"role"`
	Id       *int    `json:"id"`
}

type AccountsService struct {
	db       *sql.DB
	c_params *ConnectionParams
}

func (t *AccountsService) Insert(account Account) (int64, error) {
	var res sql.Result
	var err error

	res, err = t.db.Exec("INSERT INTO public.\"Accounts\" (username, password,role) VALUES ($1, $2)  RETURNING id", account.Username, account.Password, account.Role)

	if err != nil {
		return -1, err
	}
	var id int64
	id, err = res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (t *AccountsService) GetByID(id int) (Account, error) {

	var account Account
	err := t.db.QueryRow("SELECT username, password, role FROM public.\"Accounts\" WHERE id = $1", id).Scan(&account.Username, &account.Password, &account.Role)
	if err != nil {
		return Account{}, err
	}
	return account, nil
}

func (t *AccountsService) DeleteByID(id int) error {
	_, err := t.db.Exec("DELETE FROM public.\"Accounts\" WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
