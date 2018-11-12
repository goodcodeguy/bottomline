package processes

import (
	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/api/datastores"
	"github.com/goodcodeguy/bottomline/lib/logger"
)

var service = &ProcessConfigurationService{datastores.PrimaryDatastore, logger.New("bottomline.processes")}
var controller = &ProcessConfigurationController{service}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", service.CreateProcessConfiguration)
	router.Get("/", controller.getAllProcessConfiguration)

	router.Route("/{process_configuration_id}", func(router chi.Router) {
		router.Use(service.processConfigurationCtx)

		router.Get("/", controller.getProcessConfiguration)
		router.Put("/", controller.updateProcessConfiguration)
		router.Delete("/", controller.deleteProcessConfiguration)
	})
	return router
}
