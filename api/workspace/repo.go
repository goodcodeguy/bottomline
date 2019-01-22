package workspace

import (
	"github.com/goodcodeguy/bottomline/lib/database"
	"github.com/jinzhu/gorm"
	"github.com/juju/loggo"
)

type WorkspaceRepo struct {
	db  *database.DB
	log loggo.Logger
}

type Workspace struct {
	gorm.Model

	Name  string
	Owner int
}

func (repo WorkspaceRepo) migrate() {
	repo.db.AutoMigrate(&Workspace{})
}

func (repo WorkspaceRepo) getAllWorkspaces() []Workspace {
	workspaces := []Workspace{}
	repo.db.Find(&workspaces)
	return workspaces
}

func (repo WorkspaceRepo) getWorkspace(id int) (Workspace, error) {
	workspace := Workspace{}
	err := repo.db.Find(&workspace, id).Error
	return workspace, err
}

func (repo WorkspaceRepo) getAllWorkspacesForUser(userID int) []Workspace {
	workspaces := []Workspace{}
	repo.db.Where("owner = ?", userID).Find(&workspaces)
	return workspaces
}

func (repo WorkspaceRepo) updateWorkspace(workspace Workspace) error {
	db := repo.db.Save(workspace)
	return db.Error
}
