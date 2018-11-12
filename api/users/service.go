package users

import (
	"github.com/goodcodeguy/bottomline/api/datastores"
	"github.com/goodcodeguy/bottomline/lib/database"
	"github.com/goodcodeguy/bottomline/lib/logger"
	"github.com/juju/loggo"
)

type UserService struct {
	db  *database.DB
	log loggo.Logger
}

var Service = &UserService{datastores.PrimaryDatastore, logger.New("bottomline.processes")}
