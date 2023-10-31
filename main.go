package main

import (
	"fmt"
	"myapp/controllers"
	"myapp/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	var c_params = services.ConnectionParams{
		Host:     "localhost",
		Port:     "8080",
		User:     "",
		Password: "",
		DBname:   "SimbirGO",
	}
	db, err1 := services.СreateConnection(c_params)
	if err1 == nil {
		defer db.Close()
	} else {
		fmt.Errorf(err1.Error())
		return
	}

	accs, transp, rent, err2 := services.СreateServices(db, c_params)
	if err2 != nil {

		r := mux.NewRouter()

		var com = controllers.AccountController{Controller: controllers.Controller{func(path, method string, handler func(w http.ResponseWriter, r *http.Request)) {
			r.HandleFunc(path, handler) // Обработчик для пути "/mypath"
		}}, AService: &accs, TService: &transp, RService: &rent}
		com.SetHandlers()

		http.Handle("/", r)
		http.ListenAndServe(":3020", nil)
	}
}
