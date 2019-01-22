package processconfiguration

import (
	"github.com/goodcodeguy/bottomline/api/workspace"
	"github.com/goodcodeguy/bottomline/lib/database"
	"github.com/jinzhu/gorm"
)

type ProcessConfigurationRepo struct {
	db *database.DB
}

// ProcessConfiguration Describes the over arching configuration for a process
type ProcessConfiguration struct {
	gorm.Model

	Name          string
	Description   string
	Configuration string
	WorkspaceID   int                 `json:"-"`
	Workspace     workspace.Workspace `json:"workspace,omitempty"`
}

func (repo ProcessConfigurationRepo) migrate() {
	repo.db.AutoMigrate(&ProcessConfiguration{})
}

func (repo ProcessConfigurationRepo) getAllProcessConfigurations() []ProcessConfiguration {
	processConfigurations := []ProcessConfiguration{}
	repo.db.Preload("Workspace").Find(&processConfigurations)
	return processConfigurations
}
