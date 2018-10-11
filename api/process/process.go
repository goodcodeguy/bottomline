package process

// GetAllConfigurations Retrieves all Process Configurations
func GetAllConfigurations() []ProcessConfiguration {
	log.Infof("Querying all Process Configurations")
	qry := `select
      id,
      description,
      configuration
    from bottomline.process_configuration`

	rows, _ := db.Query(qry)

	defer rows.Close()
	log.Infof("Iterating over Rows")
	configurations := []ProcessConfiguration{}
	for rows.Next() {
		processConfiguration := ProcessConfiguration{}
		err := rows.Scan(
			&processConfiguration.ID,
			&processConfiguration.Description,
			&processConfiguration.Configuration)
		if err != nil {
			log.Criticalf("Error marshalling data from row: %s", err.Error())
		}
		configurations = append(configurations, processConfiguration)
	}
	return configurations
}

// SaveConfiguration will save a new configuration
func SaveConfiguration(conf ProcessConfiguration) {
	log.Infof("Inserting a new configuration")

	qry := `INSERT INTO bottomline.process_configuration (name, description, "configuration") VALUES ("test", "another test", "{}")`

	db.Query(qry)
}
