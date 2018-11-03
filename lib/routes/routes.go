package routes

import (
	"net/http"
)

func RouteNotImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(501), 501)
}
