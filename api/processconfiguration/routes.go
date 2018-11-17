package processconfiguration

import (
	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/api/datastores"
	"github.com/goodcodeguy/bottomline/lib/logger"
)

var service = &ProcessConfigurationService{datastores.PrimaryDatastore, logger.New("bottomline.processconfiguration")}
var controller = &ProcessConfigurationController{service}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", controller.createProcessConfiguration)
	router.Get("/", controller.getAllProcessConfiguration)

	router.Route("/{process_configuration_id}", func(router chi.Router) {
		router.Use(controller.processConfigurationCtx)

		router.Get("/", controller.getProcessConfiguration)
		router.Put("/", controller.updateProcessConfiguration)
		router.Delete("/", controller.deleteProcessConfiguration)
	})
	return router
}
