package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/api/process"
	"github.com/goodcodeguy/bottomline/lib/logger"
)

var log = logger.New("bottomline.main")

func main() {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		out, err := json.Marshal(process.GetAllConfigurations())
		if err != nil {
			log.Criticalf("Error Marshaling data to JSON: %s", err.Error())
		}
		fmt.Fprintf(w, string(out))
	})

	http.ListenAndServe(":3000", r)
}
