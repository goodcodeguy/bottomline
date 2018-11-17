package processes

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/goodcodeguy/bottomline/api/workspaces"
	"github.com/goodcodeguy/bottomline/lib/database"
	"github.com/juju/loggo"
)

type ProcessConfigurationService struct {
	db  *database.DB
	log loggo.Logger
}

// ProcessConfiguration Describes the over arching configuration for a process
type ProcessConfiguration struct {
	ID            int                  `json:"id"`
	Name          string               `json:"name"`
	Description   string               `json:"description"`
	Configuration string               `json:"configuration"`
	Workspace     workspaces.Workspace `json:"-"`
}

// GetAllConfigurations Retrieves all Process Configurations
func (svc ProcessConfigurationService) GetAllConfigurations() []ProcessConfiguration {
	svc.log.Infof("Querying all Process Configurations")
	qry := `select
      id,
			name,
      description,
      configuration
    from bottomline.process_configurations`

	rows, _ := svc.db.Query(qry)

	defer rows.Close()
	svc.log.Infof("Iterating over Rows")
	configurations := []ProcessConfiguration{}
	for rows.Next() {
		processConfiguration := ProcessConfiguration{}
		err := rows.Scan(
			&processConfiguration.ID,
			&processConfiguration.Name,
			&processConfiguration.Description,
			&processConfiguration.Configuration)
		if err != nil {
			svc.log.Criticalf("Error marshalling data from row: %s", err.Error())
		}
		configurations = append(configurations, processConfiguration)
	}
	return configurations
}

// GetProcessConfiguration allows you to get a process configuration from the database
func (svc ProcessConfigurationService) GetProcessConfiguration(id string) (ProcessConfiguration, error) {
	qry := `SELECT
						id,
						name,
						description,
						configuration
					FROM bottomline.process_configurations
					WHERE id = $1`

	p := ProcessConfiguration{}
	err := svc.db.QueryRow(qry, id).Scan(&p.ID, &p.Name, &p.Description, &p.Configuration)
	if err != nil {
		svc.log.Criticalf("Error Reading results: %s", err.Error())
	}
	return p, err
}

func (svc ProcessConfigurationService) UpdateProcessConfiguration(processConfiguration ProcessConfiguration) error {

	qry := `UPDATE bottomline.process_configurations SET name = $1, description = $2, configuration = $3 WHERE id = $4`
	err := svc.db.Exec(qry, processConfiguration.Name, processConfiguration.Description, processConfiguration.Configuration, processConfiguration.ID)

	return err

}

func (svc ProcessConfigurationService) DeleteProcessConfiguration(id string) error {
	qry := `DELETE FROM bottomline.process_configurations WHERE id = $1`
	err := svc.db.Exec(qry, id)

	return err
}

func (svc ProcessConfigurationService) CreateProcessConfiguration(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		svc.log.Criticalf("Error reading POST Body: %s", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	var p ProcessConfiguration
	err = json.Unmarshal(b, &p)
	if err != nil {
		svc.log.Criticalf("Error unmarshalling Data: %s", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	svc.log.Infof("Process Configuration: Name: %s, Description: %s, Configuration: %s", p.Name, p.Description, p.Configuration)

	err = svc.db.Exec("INSERT INTO bottomline.process_configurations (name, description, configuration) VALUES ($1, $2, $3)", p.Name, p.Description, p.Configuration)
	if err != nil {
		svc.log.Criticalf("Error Creating Process Configuration: %s", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{message: 'success'}"))
}

func (svc ProcessConfigurationService) processConfigurationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		processConfigurationID := chi.URLParam(r, "process_configuration_id")
		processConfiguration, err := svc.GetProcessConfiguration(processConfigurationID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "process_configuration", processConfiguration)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
