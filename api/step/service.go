package step

type StepService struct {
	repo *StepRepo
}

// GetAllSteps Retrieves all Process Configurations
func (svc StepService) getAllSteps() []Step {
	return svc.repo.getAllSteps()
}
