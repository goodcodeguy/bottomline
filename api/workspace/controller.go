package workspace

import (
	"encoding/json"
	"net/http"
)

type WorkspaceController struct {
	svc *WorkspaceService
}

func (ctl WorkspaceController) getAllWorkspaces(w http.ResponseWriter, r *http.Request) {
	workspaces := ctl.svc.getAllWorkspaces()

	j, err := json.Marshal(workspaces)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
