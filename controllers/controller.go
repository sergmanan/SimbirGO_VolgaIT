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

type Controller struct {
	Path_setter func(path string, method string, handler func(w http.ResponseWriter, r *http.Request))
}

func (t *Controller) createRoutes() []RouteHandler {
	return []RouteHandler{
		RouteHandler{
			route:  "/path1/path2",
			method: "GET",
			handler: func(w http.ResponseWriter, r *http.Request) {

			},
		},
		RouteHandler{
			route:  "/path1/path3",
			method: "POST",
			handler: func(w http.ResponseWriter, r *http.Request) {

			},
		},
	}
}

func (t *Controller) SetHandlers() {
	handlers := t.createRoutes()

	for _, value := range handlers {
		var handler http.Handler
		handler = http.HandlerFunc(value.handler)
		for _, middleware := range value.middlewares {
			handler = middleware(handler)
		}
		t.Path_setter(value.route, value.method, handler.ServeHTTP)
	}

}
