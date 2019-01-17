package user

import (
	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/api/datastores"
	"github.com/goodcodeguy/bottomline/lib/logger"
)

var userRepo = &UserRepo{datastores.PrimaryDatastore, logger.New("bottomline.workspace")}
var userService = &UserService{userRepo}
var userController = &UserController{userService}

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", userController.getAllUsers)

	return router
}
