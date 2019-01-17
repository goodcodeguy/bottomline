package processconfiguration

import (
	"github.com/goodcodeguy/bottomline/api/workspace"
	"github.com/goodcodeguy/bottomline/lib/database"
)

type ProcessConfigurationRepo struct {
	db *database.DB
}

// ProcessConfiguration Describes the over arching configuration for a process
type ProcessConfiguration struct {
	ID            int
	Name          string
	Description   string
	Configuration string
	Workspace     workspace.Workspace
}

func (repo ProcessConfigurationRepo) getAllProcessConfigurations() []ProcessConfiguration {
	processConfigurations := []ProcessConfiguration{}
	repo.db.Find(&processConfigurations)
	return processConfigurations
}
