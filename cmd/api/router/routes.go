package router

import (
	"github.com/gorilla/mux"
	"net/http"
	midlwr "service-account-finder/library"
	u "service-account-finder/library"
)

// InitializeRoutes it initialize all the api routes
func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", midlwr.SetMiddlewareJSON(HealthCheck)).Methods("GET")
	subrouter := router.PathPrefix("/api/saf/").Subrouter()
	subrouter.HandleFunc("/list-sa", midlwr.SetMiddlewareJSON(SAFinder)).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(u.NotFoundHandler)
	return router
}
