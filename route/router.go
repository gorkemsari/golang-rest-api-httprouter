package route

import (
	"github.com/julienschmidt/httprouter"
	"github.com/gorkemsari/golang-rest-api-httprouter/handler"
)

type Route struct {
	Name        string
	Method      string
	Path	    string
	HandlerFunc httprouter.Handle
}

type Routes []Route

var routes = Routes{
	Route{"CityAll", "GET", "/cities", handler.CityAll},
	Route{"City", "GET", "/cities/{id}", handler.City},
}

func NewRouter() *httprouter.Router {

	router := httprouter.New()
	for _, route := range routes {
		var handle httprouter.Handle

		handle = route.HandlerFunc

		router.Handle(route.Method, route.Path, handle)
	}
	return router
}