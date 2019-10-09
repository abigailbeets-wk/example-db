package project

import (
	"example-db/activity"
	"example-db/helpers"
	"example-db/workflow"
	. "example-db/workflow"
	"net/http"
)

type ProjectHandler struct{}

type Projects struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	ProjectID  int    `json:"projectID"`
	AccountID  string `json:"accountID"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Files      []File `json:"files,omitempty"`
	CreatedBy  string `json:"createdBy"`
	CreatedOn  string `json:"createdOn"`
	ModifiedBy string `json:"modifiedBy"`
	ModifiedOn string `json:"modifiedOn"`
}

type File struct {
	FileID     int     `json:"fileID"`
	CerberusID string  `json:"cerberusID"`
	Name       string  `json:"name"`
	Sheets     []Sheet `json:"sheets,omitempty"`
	CreatedBy  string  `json:"createdBy"`
	CreatedOn  string  `json:"createdOn"`
	ModifiedBy string  `json:"modifiedBy"`
	ModifiedOn string  `json:"modifiedOn"`
}

type Sheet struct {
	SheetID    int                 `json:"sheetID"`
	CerberusID string              `json:"cerberusID"`
	Name       string              `json:"name"`
	HeaderRow  int                 `json:"headerRow"`
	Headers    []Header            `json:"headers,omitempty"`
	Workflows  []workflow.Workflow `json:"workflows"`
	Revision   string              `json:"revision"`
	CreatedBy  string              `json:"createdBy"`
	CreatedOn  string              `json:"createdOn"`
	ModifiedBy string              `json:"modifiedBy"`
	ModifiedOn string              `json:"modifiedOn"`
}

type Header struct {
	HeaderID int    `json:"headerID"`
	Name     string `json:"name"`
}

func (handler *ProjectHandler) GetAllProjectsHandler(w http.ResponseWriter, r *http.Request) {
	p1 := Project{
		ProjectID: 1,
		AccountID: "1234abds",
		Name:      "My Parsing Project",
		Type:      "Parsing",
		CreatedBy: "Abby",
		CreatedOn: "10-08-2019",
	}

	p2 := Project{
		ProjectID: 2,
		AccountID: "1234abds",
		Name:      "My Parsing Project 2",
		Type:      "Parsing",
		CreatedBy: "Abby",
		CreatedOn: "10-08-2019",
	}

	projects := Projects{Projects: []Project{p1, p2}}
	helpers.RespondWithJSON(w, http.StatusOK, projects)
}

func (handler *ProjectHandler) GetProjectByIDHandler(w http.ResponseWriter, r *http.Request) {
	act := activity.Activity{
		ActivityID: 1,
		Name:       "AddPrefix",
		PrettyName: "Add Prefix",
		Index:      0,
		Hits:       0,
		CreatedBy:  "Abby",
		CreatedOn:  "10-08-2019",
		Parameters: []activity.Parameter{activity.Parameter{ParameterID: 1, Name: "Param1", Value: "Value1"}},
	}

	wf := Workflow{
		WorkflowID: 1,
		Name:       "My Workflow",
		Activities: []activity.Activity{act},
	}

	sheet := Sheet{
		SheetID:    1,
		CerberusID: "5678",
		HeaderRow:  1,
		Name:       "My Sheet",
		Revision:   "58135",
		Workflows:  []Workflow{wf},
	}

	file := File{
		FileID:     1,
		CerberusID: "1234",
		Name:       "My File",
		Sheets:     []Sheet{sheet},
	}
	project := Project{
		ProjectID: 1,
		AccountID: "1234abds",
		Name:      "My Parsing Project",
		Type:      "Parsing",
		CreatedBy: "Abby",
		CreatedOn: "10-08-2019",
		Files:     []File{file},
	}
	helpers.RespondWithJSON(w, http.StatusCreated, project)
}

func (handler *ProjectHandler) CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithStatus(w, http.StatusCreated)
}
