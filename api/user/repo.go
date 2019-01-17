package user

import (
	"github.com/goodcodeguy/bottomline/lib/database"
	"github.com/juju/loggo"
)

type User struct {
	ID   int
	Name string
}

type UserRepo struct {
	db  *database.DB
	log loggo.Logger
}

func (repo UserRepo) getAllUsers() []User {
	users := []User{}
	repo.db.Find(&users)
	return users
}

func (repo UserRepo) getUser(id int) (User, error) {
	user := User{}
	err := repo.db.Find(&user, id).Error
	return user, err
}
