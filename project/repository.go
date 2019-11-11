package project

import (
	"database/sql"
	"example-db/helpers"
	"fmt"
	"log"
	"os"
)

type ProjectRepository struct {
	DB *sql.DB
}

type IProjectRepository interface {
	GetAllProjectsForAccount(string) ([]Project, error)
}

type Project struct {
	ProjectID  int    `json:"projectID"`
	AccountID  string `json:"accountID"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	CreatedBy  string `json:"createdBy"`
	CreatedOn  string `json:"createdOn"`
	ModifiedBy string `json:"modifiedBy"`
	ModifiedOn string `json:"modifiedOn"`
}

func (repository *ProjectRepository) GetAllProjectsForAccount(accountID string) ([]Project, error) {
	query := fmt.Sprintf("SELECT ProjectID, AccountID, Name, Type, CreatedBy, CreatedDate, ModifiedBy, ModifiedDate FROM %s.Project WHERE AccountID = ?", os.Getenv("DATABASE_NAME"))

	rows, err := repository.DB.Query(query, accountID)
	if err != nil {
		errorMessage := fmt.Sprintf("error retrieving projects: %s", err.Error())
		log.Println(errorMessage)
		return nil, helpers.DBError{Code: 500, Message: errorMessage}
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		project := Project{}
		err := rows.Scan(&project.ProjectID, &project.AccountID, &project.Name, &project.Type, &project.CreatedBy, &project.CreatedOn, &project.ModifiedBy, &project.ModifiedOn)
		if err != nil {
			errorMessage := fmt.Sprintf("unable to map response to project object")
			log.Println(errorMessage)
			return nil, helpers.DBError{Code: 500, Message: errorMessage}
		}

		projects = append(projects, project)
	}

	return projects, nil
}
