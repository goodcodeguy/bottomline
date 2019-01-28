package workspace

import (
	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/api/datastores"
	"github.com/goodcodeguy/bottomline/lib/logger"
)

var workspaceRepo = &WorkspaceRepo{datastores.PrimaryDatastore, logger.New("bottomline.workspace")}
var workspaceService = &WorkspaceService{workspaceRepo}
var workspaceController = &WorkspaceController{workspaceService}

func Routes() *chi.Mux {

	router := chi.NewRouter()
	router.Post("/", workspaceController.createWorkspace)
	router.Get("/", workspaceController.getAllWorkspaces)
	router.Put("/", workspaceController.updateWorkspace)

	router.Get("/user/{user_id}", workspaceController.getWorkspacesByUserID)

	router.Route("/{workspace_id}", func(router chi.Router) {
		router.Get("/", workspaceController.getWorkspace)

		router.Delete("/", workspaceController.deleteWorkspace)
	})

	return router
}
