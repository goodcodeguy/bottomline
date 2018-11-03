package processes

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

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
