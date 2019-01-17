package user

type UserService struct {
	repo *UserRepo
}

// GetAllUsers Retrieves all Process Configurations
func (svc UserService) getAllUsers() []User {
	return svc.repo.getAllUsers()
}
