package step

import (
	"github.com/goodcodeguy/bottomline/lib/database"
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
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Status       StepStatus `json:"status"`
	ErrorMessage string     `db:"error_message" json:"error_message"`
}

func (repo StepRepo) getAllSteps() []Step {
	steps := []Step{}
	repo.db.Select(&steps, `SELECT
														id,
														name,
														description,
														status,
														error_message
													FROM bottomline.steps`)
	return steps
}
