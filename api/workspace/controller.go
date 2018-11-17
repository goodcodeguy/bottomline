package workspace

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/lib/controller"
)

type WorkspaceController struct {
	svc *WorkspaceService
}

func (ctl WorkspaceController) getAllWorkspaces(w http.ResponseWriter, r *http.Request) {
	workspaces := ctl.svc.getAllWorkspaces()

	controller.RespondWithJSON(w, workspaces)
}

func (ctl WorkspaceController) getWorkspace(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (ctl WorkspaceController) updateWorkspace(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (ctl WorkspaceController) deleteWorkspace(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (ctl WorkspaceController) createWorkspace(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (ctl WorkspaceController) workspaceCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		workspaceID := chi.URLParam(r, "workspace_id")
		workspace, err := ctl.svc.getWorkspace(workspaceID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "workspace", workspace)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
