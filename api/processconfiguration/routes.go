package processconfiguration

import (
	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/api/datastores"
	"github.com/goodcodeguy/bottomline/lib/logger"
)

var processConfigurationService = &ProcessConfigurationService{datastores.PrimaryDatastore, logger.New("bottomline.processconfiguration")}
var processConfigurationController = &ProcessConfigurationController{processConfigurationService}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", processConfigurationController.createProcessConfiguration)
	router.Get("/", processConfigurationController.getAllProcessConfiguration)

	router.Route("/{process_configuration_id}", func(router chi.Router) {
		router.Use(processConfigurationController.processConfigurationCtx)

		router.Get("/", processConfigurationController.getProcessConfiguration)
		router.Put("/", processConfigurationController.updateProcessConfiguration)
		router.Delete("/", processConfigurationController.deleteProcessConfiguration)
	})
	return router
}
