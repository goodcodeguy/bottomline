package workspaces

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", getAllWorkspaces)

	return router
}

func getAllWorkspaces(w http.ResponseWriter, r *http.Request) {
	workspaces := GetAllWorkspaces()

	j, err := json.Marshal(workspaces)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
