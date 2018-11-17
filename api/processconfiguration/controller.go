package processconfiguration

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

type ProcessConfigurationController struct {
	svc *ProcessConfigurationService
}

func (ctl ProcessConfigurationController) getAllProcessConfiguration(w http.ResponseWriter, r *http.Request) {
	processConfigurations := ctl.svc.GetAllConfigurations()

	j, err := json.Marshal(processConfigurations)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func (ctl ProcessConfigurationController) getProcessConfiguration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	processConfiguration, ok := ctx.Value("process_configuration").(ProcessConfiguration)
	if !ok {
		ctl.svc.log.Criticalf("Error when sending response")
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	j, err := json.Marshal(processConfiguration)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func (ctl ProcessConfigurationController) deleteProcessConfiguration(w http.ResponseWriter, r *http.Request) {
	processConfigurationID := chi.URLParam(r, "process_configuration_id")
	err := ctl.svc.DeleteProcessConfiguration(processConfigurationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.StatusText(204)
}

func (ctl ProcessConfigurationController) updateProcessConfiguration(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		ctl.svc.log.Criticalf("Error reading POST Body: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var p ProcessConfiguration
	err = json.Unmarshal(b, &p)
	if err != nil {
		ctl.svc.log.Criticalf("Error unmarshalling Data: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ctl.svc.UpdateProcessConfiguration(p)
	if err != nil {
		ctl.svc.log.Criticalf("Error updating Process Configuration: %s", err.Error())
	}

	http.StatusText(http.StatusNoContent)
}
