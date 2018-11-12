package users

// GetAllUsers Retrieves all Process Configurations
func (svc UserService) GetAllUsers() []User {
	svc.log.Infof("Querying all Users")
	qry := `select
      id,
			name
    from bottomline.users`

	rows, _ := svc.db.Query(qry)

	defer rows.Close()
	svc.log.Infof("Iterating over Rows")
	users := []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(
			&user.ID,
			&user.Name)
		if err != nil {
			svc.log.Criticalf("Error marshalling data from row: %s", err.Error())
		}
		users = append(users, user)
	}
	return users
}
