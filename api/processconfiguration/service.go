package processconfiguration

import (
	"github.com/goodcodeguy/bottomline/api/workspace"
	"github.com/goodcodeguy/bottomline/lib/database"
	"github.com/juju/loggo"
)

type ProcessConfigurationService struct {
	db  *database.DB
	log loggo.Logger
}

// ProcessConfiguration Describes the over arching configuration for a process
type ProcessConfiguration struct {
	ID            int                 `json:"id"`
	Name          string              `json:"name"`
	Description   string              `json:"description"`
	Configuration string              `json:"configuration"`
	Workspace     workspace.Workspace `json:"-"`
}

func (svc ProcessConfigurationService) getAllConfigurations() []ProcessConfiguration {
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
func (svc ProcessConfigurationService) getProcessConfiguration(id string) (ProcessConfiguration, error) {
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

func (svc ProcessConfigurationService) updateProcessConfiguration(processConfiguration ProcessConfiguration) error {

	qry := `UPDATE bottomline.process_configurations SET name = $1, description = $2, configuration = $3 WHERE id = $4`
	err := svc.db.Exec(qry, processConfiguration.Name, processConfiguration.Description, processConfiguration.Configuration, processConfiguration.ID)

	return err

}

func (svc ProcessConfigurationService) deleteProcessConfiguration(id string) error {
	qry := `DELETE FROM bottomline.process_configurations WHERE id = $1`
	err := svc.db.Exec(qry, id)

	return err
}

func (svc ProcessConfigurationService) createProcessConfiguration(processConfiguration ProcessConfiguration) error {
	qry := `INSERT INTO bottomline.process_configurations (name, description, configuration) VALUES ($1, $2, $3)`
	err := svc.db.Exec(qry, processConfiguration.Name, processConfiguration.Description, processConfiguration.Configuration)

	return err
}
