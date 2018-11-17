package user

import (
	"net/http"

	"github.com/goodcodeguy/bottomline/lib/controller"
)

type UserController struct {
	svc *UserService
}

func (ctl UserController) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := ctl.svc.getAllUsers()

	controller.RespondWithJSON(w, users)
}
