package workspace

import (
	"net/http"

	"github.com/goodcodeguy/bottomline/lib/controller"
)

type WorkspaceController struct {
	svc *WorkspaceService
}

func (ctl WorkspaceController) getAllWorkspaces(w http.ResponseWriter, r *http.Request) {
	workspaces := ctl.svc.getAllWorkspaces()

	controller.RespondWithJSON(w, workspaces)
}
