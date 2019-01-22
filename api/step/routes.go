package step

import (
	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/api/datastores"
	"github.com/goodcodeguy/bottomline/lib/logger"
)

var stepRepo = &StepRepo{datastores.PrimaryDatastore, logger.New("bottomline.step")}
var stepService = &StepService{stepRepo}
var stepController = &StepController{stepService}

func migrateModels() {
	stepRepo.migrate()
}

func Routes() *chi.Mux {
	migrateModels()

	router := chi.NewRouter()

	router.Get("/", stepController.getAllSteps)

	return router
}
