package step

import (
	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/api/datastores"
	"github.com/goodcodeguy/bottomline/lib/logger"
)

var service = &StepService{datastores.PrimaryDatastore, logger.New("bottomline.step")}
var controller = &StepController{service}

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", controller.getAllSteps)

	return router
}
