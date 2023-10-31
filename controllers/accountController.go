package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"myapp/services"
	"net/http"
)

type AccountController struct {
	Controller
	AService *services.AccountsService
	TService *services.TransportService
	RService *services.RentService
}

func (t *AccountController) CreateRoutes() []RouteHandler {
	return []RouteHandler{
		RouteHandler{
			route:  "/Account/Me",
			method: "GET",
			handler: func(w http.ResponseWriter, r *http.Request) {
				//описание: получение данных о текущем аккаунте
				//ограничения: только авторизованные пользователи
				fmt.Fprintln(w, "Привет, мир!")
			},
		},
		RouteHandler{
			route:  "/Account/SignUp",
			method: "POST",
			handler: func(w http.ResponseWriter, r *http.Request) {
				var acc services.Account
				decoder := json.NewDecoder(r.Body)

				if err := decoder.Decode(&acc); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "Ошибка при разборе JSON: %v", err)
					log.Println(err)
					return
				}

				id, err := t.AService.Insert(acc)

				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "Ошибка при вставке: %v", err)
					log.Println(err)
					return
				}

				res := struct{ inserted_id int64 }{
					inserted_id: id,
				}

				// Преобразуем структуру в JSON
				jsonData, err := json.Marshal(res)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, "Ошибка при создании JSON: %v", err)
					log.Println(err)
					return
				}

				// Устанавливаем заголовок Content-Type на application/json
				w.Header().Set("Content-Type", "application/json")

				// Отправляем JSON в ответе HTTP
				w.WriteHeader(http.StatusOK)
				w.Write(jsonData)

			},
		},
		RouteHandler{
			route:  "/Account/SignIn",
			method: "POST",
			handler: func(w http.ResponseWriter, r *http.Request) {
				// 				описание: получение нового jwt токена пользователя
				// body:
				// {
				// "username": "string", //имя пользователя
				// "password": "string" //пароль
				// }
				// ограничения: нет
			},
		},
		RouteHandler{
			route:  "/Account/SignOut",
			method: "POST",
			handler: func(w http.ResponseWriter, r *http.Request) {

			},
		},
		RouteHandler{
			route:  "/Account/Update",
			method: "PUT",
			handler: func(w http.ResponseWriter, r *http.Request) {
				// 				описание: обновление своего аккаунта
				// body:
				// {
				// "username": "string", //имя пользователя
				// "password": "string" //пароль
				// }
				// ограничения: только авторизованные пользователи, нельзя использовать уже
				// используемые в системе username
			},
		},
	}
}
