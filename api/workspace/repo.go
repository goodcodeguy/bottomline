package workspace

import (
	"github.com/goodcodeguy/bottomline/lib/database"
	"github.com/juju/loggo"
)

type WorkspaceRepo struct {
	db  *database.DB
	log loggo.Logger
}

type Workspace struct {
	Name  string
	Owner int
}

func (repo WorkspaceRepo) getAllWorkspaces() []Workspace {
	workspaces := []Workspace{}

	return workspaces
}

func (repo WorkspaceRepo) getWorkspace(id int) (Workspace, error) {
	workspace := Workspace{}

	return workspace, nil
}

func (repo WorkspaceRepo) getAllWorkspacesForUser(userID int) []Workspace {
	workspaces := []Workspace{}

	return workspaces
}

func (repo WorkspaceRepo) updateWorkspace(workspace Workspace) error {

	return nil
}
