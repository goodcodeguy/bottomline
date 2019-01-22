package workspace

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

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
	strWorkspaceID := chi.URLParam(r, "workspace_id")
	workspaceID, err := strconv.Atoi(strWorkspaceID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	workspace, err := ctl.svc.getWorkspace(workspaceID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	controller.RespondWithJSON(w, workspace)
}

func (ctl WorkspaceController) getWorkspacesByUserID(w http.ResponseWriter, r *http.Request) {
	strUserID := chi.URLParam(r, "user_id")
	userID, err := strconv.Atoi(strUserID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	workspaces := ctl.svc.getAllWorkspacesForUser(userID)

	controller.RespondWithJSON(w, workspaces)
}

func (ctl WorkspaceController) updateWorkspace(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var workspace Workspace
	err = json.Unmarshal(b, &workspace)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctl.svc.updateWorkspace(workspace)
	controller.RespondWithJSON(w, workspace)
}

func (ctl WorkspaceController) deleteWorkspace(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (ctl WorkspaceController) createWorkspace(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
