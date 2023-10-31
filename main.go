package main

import (
	"myapp/services"
)

func main() {
	var err error
	err.Error()
	var c_params = services.ConnectionParams{
		Host:     "localhost",
		Port:     "8080",
		User:     "",
		Password: "",
		DBname:   "SimbirGO",
	}
	db, err := services.СreateConnection(c_params)

	accs, transp, rent, err := services.СreateServices(db, c_params)

}
