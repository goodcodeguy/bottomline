package users

import (
	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/lib/routes"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", routes.RouteNotImplemented)

	return router
}
