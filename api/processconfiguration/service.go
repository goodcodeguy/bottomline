package processconfiguration

import "github.com/juju/loggo"

type ProcessConfigurationService struct {
	repo *ProcessConfigurationRepo
	log  loggo.Logger
}

func (svc ProcessConfigurationService) getAllConfigurations() []ProcessConfiguration {
	return svc.repo.getAllProcessConfigurations()
}

// GetProcessConfiguration allows you to get a process configuration from the database
func (svc ProcessConfigurationService) getProcessConfiguration(id string) ProcessConfiguration {
	p := ProcessConfiguration{}

	return p
}

func (svc ProcessConfigurationService) updateProcessConfiguration(processConfiguration ProcessConfiguration) {

}

func (svc ProcessConfigurationService) deleteProcessConfiguration(id string) {

}

func (svc ProcessConfigurationService) createProcessConfiguration(processConfiguration ProcessConfiguration) {

}
