package user

import (
	"net/http"

	"github.com/goodcodeguy/bottomline/lib/controller"
)

type UserController struct {
	userService *UserService
}

func (ctl UserController) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := ctl.userService.getAllUsers()

	controller.RespondWithJSON(w, users)
}

func (ctl UserController) getAllUserWorkspaces(w http.ResponseWriter, r *http.Request) {

}
