package process

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", CreateProcessConfiguration)
	router.Get("/", getAllProcessConfiguration)

	router.Route("/{processConfigurationID}", func(router chi.Router) {
		router.Use(processConfigurationCtx)

		router.Get("/", getProcessConfiguration)
		router.Put("/", updateProcessConfiguration)
		router.Delete("/", deleteProcessConfiguration)
	})
	return router
}

func processConfigurationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		processConfigurationID := chi.URLParam(r, "processConfigurationID")
		processConfiguration, err := GetProcessConfiguration(processConfigurationID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "processConfiguration", processConfiguration)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllProcessConfiguration(w http.ResponseWriter, r *http.Request) {
	processConfigurations := GetAllConfigurations()

	j, err := json.Marshal(processConfigurations)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func getProcessConfiguration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	processConfiguration, ok := ctx.Value("processConfiguration").(ProcessConfiguration)
	if !ok {
		log.Criticalf("Error when sending response")
		http.Error(w, http.StatusText(422), 422)
		return
	}

	j, err := json.Marshal(processConfiguration)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func deleteProcessConfiguration(w http.ResponseWriter, r *http.Request) {
	processConfigurationID := chi.URLParam(r, "processConfigurationID")
	err := DeleteProcessConfiguration(processConfigurationID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.StatusText(204)
}

func updateProcessConfiguration(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Criticalf("Error reading POST Body: %s", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	var p ProcessConfiguration
	err = json.Unmarshal(b, &p)
	if err != nil {
		log.Criticalf("Error unmarshalling Data: %s", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	err = UpdateProcessConfiguration(p)
	if err != nil {
		log.Criticalf("Error updating Process Configuration: %s", err.Error())
	}

	http.StatusText(204)
}
