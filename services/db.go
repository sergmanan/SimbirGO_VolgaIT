package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type ConnectionParams struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
}

func СloseConnection(db *sql.DB) error {
	if db == nil {
		return nil
	}

	err := db.Close()

	if err != nil {
		return err
	}

	return nil
}

func СreateConnection(c_params ConnectionParams) (*sql.DB, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c_params.Host, c_params.Port, c_params.User, c_params.Password, c_params.DBname)
	db, err := sql.Open("postgres", connStr)
	//db.SetConnMaxLifetime(time.Second * 4)
	if err != nil {
		log.Println(err, "init")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err, "connect")
		panic(err)
	}

	return db, nil
}

func СreateServices(db *sql.DB, c_params ConnectionParams) (a AccountsService, t TransportService, r RentService, e error) {
	r = RentService{db, &c_params}
	t = TransportService{db, &c_params}
	a = AccountsService{db, &c_params}
	e = nil
	return a, t, r, e
}
