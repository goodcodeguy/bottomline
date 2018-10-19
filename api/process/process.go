package process

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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

// GetProcessConfiguration allows you to get a process configuration from the database
func GetProcessConfiguration(id string) (ProcessConfiguration, error) {
	qry := `SELECT
						id,
						name,
						description,
						configuration
					FROM bottomline.process_configuration
					WHERE id = $1`

	p := ProcessConfiguration{}
	err := db.QueryRow(qry, id).Scan(&p.ID, &p.Name, &p.Description, &p.Configuration)
	if err != nil {
		log.Criticalf("Error Reading results: %s", err.Error())
	}
	return p, err
}

func UpdateProcessConfiguration(processConfiguration ProcessConfiguration) error {

	qry := `UPDATE bottomline.process_configuration SET name = $1, description = $2, configuration = $3 WHERE id = $4`
	err := db.Exec(qry, processConfiguration.Name, processConfiguration.Description, processConfiguration.Configuration, processConfiguration.ID)

	return err

}

func DeleteProcessConfiguration(id string) error {
	qry := `DELETE FROM bottomline.process_configuration WHERE id = $1`
	err := db.Exec(qry, id)

	return err
}

func CreateProcessConfiguration(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Criticalf("Error reading POST Body: %s", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	var p ProcessConfiguration
	err = json.Unmarshal(b, &p)
	if err != nil {
		log.Criticalf("Error unmarshalling Data: %s", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	log.Infof("Process Configuration: Name: %s, Description: %s, Configuration: %s", p.Name, p.Description, p.Configuration)

	err = db.Exec("INSERT INTO bottomline.process_configuration (name, description, configuration) VALUES ($1, $2, $3)", p.Name, p.Description, p.Configuration)
	if err != nil {
		log.Criticalf("Error Creating Process Configuration: %s", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{message: 'success'}"))
}
