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
	Name          string
	Description   string
	Configuration string
	WorkspaceID   int                 `json:"-"`
	Workspace     workspace.Workspace `json:"workspace,omitempty"`
}

func (repo ProcessConfigurationRepo) getAllProcessConfigurations() []ProcessConfiguration {
	processConfigurations := []ProcessConfiguration{}

	return processConfigurations
}
