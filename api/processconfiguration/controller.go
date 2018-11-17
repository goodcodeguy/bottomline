package processconfiguration

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

type ProcessConfigurationController struct {
	svc *ProcessConfigurationService
}

func (ctl ProcessConfigurationController) getAllProcessConfiguration(w http.ResponseWriter, r *http.Request) {
	processConfigurations := ctl.svc.getAllConfigurations()

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
	err := ctl.svc.deleteProcessConfiguration(processConfigurationID)
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

	err = ctl.svc.updateProcessConfiguration(p)
	if err != nil {
		ctl.svc.log.Criticalf("Error updating Process Configuration: %s", err.Error())
	}

	http.StatusText(http.StatusNoContent)
}

func (ctl ProcessConfigurationController) createProcessConfiguration(w http.ResponseWriter, r *http.Request) {

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

	ctl.svc.log.Infof("Process Configuration: Name: %s, Description: %s, Configuration: %s", p.Name, p.Description, p.Configuration)

	err = ctl.svc.createProcessConfiguration(p)
	if err != nil {
		ctl.svc.log.Criticalf("Error Creating Process Configuration: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{message: 'success'}"))
}

func (ctl ProcessConfigurationController) processConfigurationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		processConfigurationID := chi.URLParam(r, "process_configuration_id")
		processConfiguration, err := ctl.svc.getProcessConfiguration(processConfigurationID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "process_configuration", processConfiguration)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
