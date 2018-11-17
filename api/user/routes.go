package user

import (
	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/api/datastores"
	"github.com/goodcodeguy/bottomline/lib/logger"
)

var service = &UserService{datastores.PrimaryDatastore, logger.New("bottomline.processes")}
var controller = &UserController{service}

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", controller.getAllUsers)

	return router
}
