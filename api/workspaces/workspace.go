package workspaces

// GetAllWorkspaces Retrieves all Workspaces
func GetAllWorkspaces() []Workspace {
	log.Infof("Querying all Process Configurations")
	qry := `select
      id,
			name
    from bottomline.workspaces`

	rows, _ := db.Query(qry)

	defer rows.Close()
	log.Infof("Iterating over Rows")
	workspaces := []Workspace{}
	for rows.Next() {
		workspace := Workspace{}
		err := rows.Scan(
			&workspace.ID,
			&workspace.Name)
		if err != nil {
			log.Criticalf("Error marshalling data from row: %s", err.Error())
		}
		workspaces = append(workspaces, workspace)
	}
	return workspaces
}
