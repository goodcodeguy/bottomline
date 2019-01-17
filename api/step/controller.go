package step

import (
	"net/http"

	"github.com/goodcodeguy/bottomline/lib/controller"
)

type StepController struct {
	svc *StepService
}

func (ctl StepController) getAllSteps(w http.ResponseWriter, r *http.Request) {
	users := ctl.svc.getAllSteps()

	controller.RespondWithJSON(w, users)
}
