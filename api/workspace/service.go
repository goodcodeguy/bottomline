package workspace

import (
	"github.com/goodcodeguy/bottomline/lib/database"
	"github.com/juju/loggo"
)

type WorkspaceService struct {
	db  *database.DB
	log loggo.Logger
}

type Workspace struct {
	ID   int
	Name string
}

// GetAllWorkspaces Retrieves all Workspaces
func (svc WorkspaceService) getAllWorkspaces() []Workspace {
	svc.log.Infof("Querying all Process Configurations")
	qry := `select
      id,
			name
    from bottomline.workspaces`

	rows, _ := svc.db.Query(qry)

	defer rows.Close()
	svc.log.Infof("Iterating over Rows")
	workspaces := []Workspace{}
	for rows.Next() {
		workspace := Workspace{}
		err := rows.Scan(
			&workspace.ID,
			&workspace.Name)
		if err != nil {
			svc.log.Criticalf("Error marshalling data from row: %s", err.Error())
		}
		workspaces = append(workspaces, workspace)
	}
	return workspaces
}
