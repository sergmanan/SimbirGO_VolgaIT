package controllers

import (
	"fmt"
	"myapp/services"
	"net/http"
)

type AccountController struct {
	Controller
	AService *services.AccountsService
	TService *services.TransportService
	RService *services.RentService
}

func (t *AccountController) createRoutes() []RouteHandler {
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
