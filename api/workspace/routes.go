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
	router.Post("/", workspaceController.createWorkspace)
	router.Get("/", workspaceController.getAllWorkspaces)

	router.Route("/{workspace_id}", func(router chi.Router) {
		router.Use(workspaceController.workspaceCtx)

		router.Get("/", workspaceController.getWorkspace)
		router.Put("/", workspaceController.updateWorkspace)
		router.Delete("/", workspaceController.deleteWorkspace)
	})

	return router
}
