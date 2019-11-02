package platform

import "github.com/gorilla/mux"

// NewRouter - Provides a constructor to inject all the routable capable objects.
func NewRouter(routables ...Routable) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, routeProvider := range routables {
		routes := routeProvider.Routes()
		for _, route := range routes {
			handler := Logger(route.handler, route.name)
			router.
				Methods(route.method).
				Path(route.pattern).
				Name(route.name).
				Handler(handler)
		}
	}
	return router
}
