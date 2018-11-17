package user

import (
	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/api/datastores"
	"github.com/goodcodeguy/bottomline/lib/logger"
)

var userService = &UserService{datastores.PrimaryDatastore, logger.New("bottomline.user")}
var userController = &UserController{userService}

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", userController.getAllUsers)

	return router
}
