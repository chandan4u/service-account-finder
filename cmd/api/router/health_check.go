package router

import (
	"net/http"
	u "service-account-finder/library"
)

/*
	[ http call for healthcheck ]
	[ GET ] http://localhost:8080/healthcheck
*/

// HealthCheck api endpoint for checking service health
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, u.Message("OK", "Health check OK"), 200)
	return
}
