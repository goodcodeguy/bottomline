package step

import (
	"net/http"
)

type StepController struct {
	svc *StepService
}

func (ctl StepController) getAllSteps(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
