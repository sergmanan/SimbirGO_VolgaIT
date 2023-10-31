package main

import (
	"fmt"
	"log"
	"myapp/controllers"
	"myapp/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	var c_params = services.ConnectionParams{
		Host:     "localhost",
		Port:     "8000",
		User:     "postgres",
		Password: "ttt",
		DBname:   "SimbirGO",
	}
	db, err1 := services.СreateConnection(c_params)
	if err1 == nil {
		//defer db.Close()
	} else {
		log.Println(err1)
		return
	}

	accs, transp, rent, err2 := services.СreateServices(db, c_params)
	if err2 == nil {

		r := mux.NewRouter()

		var com = controllers.AccountController{AService: &accs, TService: &transp, RService: &rent}
		controllers.SetHandlers(func(path, method string, handler func(w http.ResponseWriter, r *http.Request)) {
			r.HandleFunc(path, handler) // Обработчик для пути "/mypath"
			fmt.Println(path)
		}, com.CreateRoutes())

		http.Handle("/", r)
		http.ListenAndServe(":3020", nil)
	} else {
		log.Println(err2)
		return
	}
}
