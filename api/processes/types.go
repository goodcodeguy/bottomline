package processes

import "github.com/goodcodeguy/bottomline/api/workspaces"

// RunningInstance is a representation of a running process
type RunningInstance struct {
	ID   int
	Name string
}

// ProcessConfiguration Describes the over arching configuration for a process
type ProcessConfiguration struct {
	ID            int                  `json:"id"`
	Name          string               `json:"name"`
	Description   string               `json:"description"`
	Configuration string               `json:"configuration"`
	Workspace     workspaces.Workspace `json:"-"`
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
	ID           int
	Name         string
	Description  string
	Status       StepStatus
	ErrorMessage string
}
