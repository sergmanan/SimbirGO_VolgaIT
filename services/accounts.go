package services

import (
	"database/sql"
	"errors"
	"log"
	"strings"
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

	// Подготовка SQL-запроса с возвращением ID (serial primary key)
	stmt, err := t.db.Prepare("INSERT INTO public.\"Accounts\" (username, password) VALUES ($1, $2)  RETURNING id")
	if err != nil {
		log.Println(err, "ins", t.db.Ping())
		return -1, err
	}

	var id int64
	err = stmt.QueryRow(account.Username, account.Password).Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "Accounts_username_key") {
			log.Println(err, "getval")
			return -1, errors.New("this username already exist")
		}
		log.Println(err, "getval")
		return -1, errors.New("database error")
	}

	return id, nil
}

func (t *AccountsService) GetByID(id int) (Account, error) {

	var account Account
	err := t.db.QueryRow("SELECT username, password, role FROM public.\"Accounts\" WHERE id = $1", id).Scan(&account.Username, &account.Password, &account.Role)
	if err != nil {
		log.Println(err)
		return Account{}, err
	}
	return account, nil
}

func (t *AccountsService) DeleteByID(id int) error {
	_, err := t.db.Exec("DELETE FROM public.\"Accounts\" WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
