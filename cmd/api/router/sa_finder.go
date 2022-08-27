package router

import (
	"net/http"
	"service-account-finder/internal/safinder"
	u "service-account-finder/library"
)

func SAFinder(resW http.ResponseWriter, reqR *http.Request) {
	safinder.SAFinder()
	u.Respond(resW, u.Message("OK", "SA Finder"), 200)
	return
}
