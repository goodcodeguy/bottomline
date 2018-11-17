package workspace

import (
	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/api/datastores"
	"github.com/goodcodeguy/bottomline/lib/logger"
)

var service = &WorkspaceService{datastores.PrimaryDatastore, logger.New("bottomline.workspace")}
var controller = &WorkspaceController{service}

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", controller.getAllWorkspaces)

	return router
}
