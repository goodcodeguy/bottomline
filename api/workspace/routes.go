package workspace

import (
	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/api/datastores"
	"github.com/goodcodeguy/bottomline/lib/logger"
)

var workspaceService = &WorkspaceService{datastores.PrimaryDatastore, logger.New("bottomline.workspace")}
var workspaceController = &WorkspaceController{workspaceService}

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", workspaceController.getAllWorkspaces)

	return router
}
