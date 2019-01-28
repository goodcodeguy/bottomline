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
	database.Model

	Name    string `json:"name"`
	OwnerId int    `db:"owner_id" json:"owner_id"`
}

func (repo WorkspaceRepo) getAllWorkspaces() []Workspace {
	workspaces := []Workspace{}
	repo.db.Select(&workspaces, `SELECT
																	id,
																	name,
																	owner_id,
																	created_at,
																	updated_at
															 FROM bottomline.workspaces`)
	return workspaces
}

func (repo WorkspaceRepo) getWorkspace(id int) (Workspace, error) {
	stmt, err := repo.db.Preparex(`SELECT
																		id,
																		name,
																		owner_id,
																		created_at,
																		updated_at
																	FROM bottomline.workspaces
																	WHERE id = $1`)
	if err != nil {
		panic("Error when preparing statement")
	}

	workspace := Workspace{}
	err = stmt.Get(&workspace, id)

	return workspace, err
}

func (repo WorkspaceRepo) getAllWorkspacesForUser(userID int) []Workspace {
	workspaces := []Workspace{}
	repo.db.Select(&workspaces, `SELECT
																	id,
																	name,
																	owner_id,
																	created_at,
																	updated_at
															 FROM bottomline.workspaces
															 WHERE owner_id = $1`, userID)
	return workspaces
}

func (repo WorkspaceRepo) updateWorkspace(workspace Workspace) error {

	return nil
}
