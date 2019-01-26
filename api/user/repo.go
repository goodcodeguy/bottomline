package user

import (
	"github.com/goodcodeguy/bottomline/lib/database"
	"github.com/juju/loggo"
)

type User struct {
	Name string
}

type UserRepo struct {
	db  *database.DB
	log loggo.Logger
}

func (repo UserRepo) getAllUsers() []User {
	users := []User{}

	return users
}

func (repo UserRepo) getUser(id int) (User, error) {
	user := User{}

	return user, nil
}
