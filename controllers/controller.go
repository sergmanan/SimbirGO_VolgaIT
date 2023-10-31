package controllers

import "net/http"

type RouteHandler struct {
	route       string
	method      string
	middlewares []func(http.Handler) http.Handler
	handler     func(http.ResponseWriter, *http.Request)
}

func createMiddleware(callback func(http.ResponseWriter, *http.Request, http.Handler)) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			callback(w, r, next)
		})
	}
}

type Controller interface {
	CreateRoutes() []RouteHandler
}

func SetHandlers(path_setter func(path string, method string, handler func(w http.ResponseWriter, r *http.Request)), handlers []RouteHandler) {

	for _, value := range handlers {
		var handler http.Handler
		handler = http.HandlerFunc(value.handler)
		for _, middleware := range value.middlewares {
			handler = middleware(handler)
		}
		path_setter(value.route, value.method, handler.ServeHTTP)
	}

}
