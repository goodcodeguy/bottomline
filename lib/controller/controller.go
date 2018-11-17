package controller

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, object interface{}) {
	j, err := json.Marshal(object)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
