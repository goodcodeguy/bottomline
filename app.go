package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/api/process"
	"github.com/goodcodeguy/bottomline/config"
	"github.com/goodcodeguy/bottomline/lib/logger"
)

var cfg = config.GetConfiguration()
var log = logger.New("bottomline.main")

func main() {

	r := chi.NewRouter()

	r.Route("/processConfiguration", func(r chi.Router) {
		r.Post("/", process.CreateProcessConfiguration)

		r.Route("/{processConfigurationID}", func(r chi.Router) {
			r.Use(processConfigurationCtx)

			r.Get("/", getProcessConfiguration)
			r.Put("/", updateProcessConfiguration)
			r.Delete("/", deleteProcessConfiguration)
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		out, err := json.Marshal(process.GetAllConfigurations())
		if err != nil {
			log.Criticalf("Error Marshaling data to JSON: %s", err.Error())
		}
		fmt.Fprintf(w, string(out))
	})

	log.Infof("Starting server on port %s", cfg.ServicePort)

	http.ListenAndServe(":"+cfg.ServicePort, r)
}

func processConfigurationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		processConfigurationID := chi.URLParam(r, "processConfigurationID")
		processConfiguration, err := process.GetProcessConfiguration(processConfigurationID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "processConfiguration", processConfiguration)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getProcessConfiguration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	processConfiguration, ok := ctx.Value("processConfiguration").(process.ProcessConfiguration)
	if !ok {
		log.Criticalf("Error when sending response")
		http.Error(w, http.StatusText(422), 422)
		return
	}

	w.Write([]byte(fmt.Sprintf("id: %d", processConfiguration.ID)))
}

func deleteProcessConfiguration(w http.ResponseWriter, r *http.Request) {
	processConfigurationID := chi.URLParam(r, "processConfigurationID")
	err := process.DeleteProcessConfiguration(processConfigurationID)
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

	var p process.ProcessConfiguration
	err = json.Unmarshal(b, &p)
	if err != nil {
		log.Criticalf("Error unmarshalling Data: %s", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	err = process.UpdateProcessConfiguration(p)
	if err != nil {
		log.Criticalf("Error updating Process Configuration: %s", err.Error())
	}

	http.StatusText(204)
}
