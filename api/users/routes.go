package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", getAllUsers)

	return router
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := Service.GetAllUsers()

	j, err := json.Marshal(users)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
