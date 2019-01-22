package workspace

type WorkspaceService struct {
	repo *WorkspaceRepo
}

// GetAllWorkspaces Retrieves all Workspaces
func (svc WorkspaceService) getAllWorkspaces() []Workspace {
	return svc.repo.getAllWorkspaces()
}

func (svc WorkspaceService) getWorkspace(id int) (Workspace, error) {
	return svc.repo.getWorkspace(id)
}

func (svc WorkspaceService) getAllWorkspacesForUser(userID int) []Workspace {
	return svc.repo.getAllWorkspacesForUser(userID)
}

func (svc WorkspaceService) updateWorkspace(workspace Workspace) {
	svc.repo.updateWorkspace(workspace)
}
