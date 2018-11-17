package user

import (
	"encoding/json"
	"net/http"
)

type UserController struct {
	svc *UserService
}

func (ctl UserController) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := ctl.svc.getAllUsers()

	j, err := json.Marshal(users)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
