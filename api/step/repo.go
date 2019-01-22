package step

import (
	"github.com/goodcodeguy/bottomline/lib/database"
	"github.com/jinzhu/gorm"
	"github.com/juju/loggo"
)

type StepRepo struct {
	db  *database.DB
	log loggo.Logger
}

// StepStatus describes the status of a step
type StepStatus int

const (
	NotStarted StepStatus = 0
	InProgress StepStatus = 1
	Complete   StepStatus = 3
	Error      StepStatus = -1
)

func (s StepStatus) String() string {
	v := [...]string{
		"NotStarted",
		"InProgress",
		"Complete",
		"Error",
	}
	return v[s]
}

type Step struct {
	gorm.Model

	Name         string
	Description  string
	Status       StepStatus
	ErrorMessage string
}

func (repo StepRepo) migrate() {
	repo.db.AutoMigrate(&Step{})
}

func (repo StepRepo) getAllSteps() []Step {
	steps := []Step{}
	repo.db.Find(&steps)
	return steps
}
