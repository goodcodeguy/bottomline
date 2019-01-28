package user

import (
	"github.com/goodcodeguy/bottomline/lib/database"
	"github.com/juju/loggo"
)

type User struct {
	database.Model

	Name string `json:"name"`
}

type UserRepo struct {
	db  *database.DB
	log loggo.Logger
}

func (repo UserRepo) getAllUsers() []User {
	users := []User{}
	repo.db.Select(&users, "SELECT id, name, created_at, updated_at FROM bottomline.users")
	return users
}

func (repo UserRepo) getUser(id int) (User, error) {
	user := User{}
	err := repo.db.Get(&user, "SELECT * FROM bottomline.users WHERE id = ?", id)
	return user, err
}
