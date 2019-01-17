package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/goodcodeguy/bottomline/api/processconfiguration"
	"github.com/goodcodeguy/bottomline/api/step"
	"github.com/goodcodeguy/bottomline/api/user"
	"github.com/goodcodeguy/bottomline/api/workspace"
	"github.com/goodcodeguy/bottomline/config"
	"github.com/goodcodeguy/bottomline/lib/logger"
)

var cfg = config.GetConfiguration()
var log = logger.New("bottomline.main")

// Routes define the primary routes for the webservice
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/workspaces", workspace.Routes())
		r.Mount("/process-configurations", processconfiguration.Routes())
		r.Mount("/steps", step.Routes())
		r.Mount("/users", user.Routes())
	})

	return router
}

func main() {
	router := Routes()

	log.Infof("Starting server on port %s", cfg.ServicePort)
	log.Infof("Available Routes")

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Infof("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Criticalf("Logging err: %s\n", err.Error())
	}

	http.ListenAndServe(":"+cfg.ServicePort, router)
}
