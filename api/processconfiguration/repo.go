package processconfiguration

import (
	"github.com/goodcodeguy/bottomline/lib/database"
)

type ProcessConfigurationRepo struct {
	db *database.DB
}

// ProcessConfiguration Describes the over arching configuration for a process
type ProcessConfiguration struct {
	database.Model

	Name          string `json:"name"`
	Description   string `json:"description"`
	Configuration string `json:"configuration"`
	WorkspaceID   int    `db:"workspace_id" json:"-"`
}

func (repo ProcessConfigurationRepo) getAllProcessConfigurations() []ProcessConfiguration {
	processConfigurations := []ProcessConfiguration{}
	repo.db.Select(&processConfigurations, `SELECT
																						id,
																						name,
																						description,
																						configuration,
																						workspace_id,
																						created_at,
																						updated_at
																					FROM bottomline.process_configurations`)
	return processConfigurations
}
