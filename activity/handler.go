package activity

import (
	"example-db/helpers"
	"net/http"
)

type ActivityHandler struct{}

type Activities struct {
	Activities []Activity `json:"activities"`
}

type Activity struct {
	ActivityID int         `json:"activityID"`
	Name       string      `json:"name"`
	PrettyName string      `json:"prettyName"`
	Index      int         `json:"index"`
	Note       string      `json:"note"`
	Summary    string      `json:"summary"`
	Hits       int         `json:"hits"`
	Parameters []Parameter `json:"parameters,omitempty"`
	CreatedBy  string      `json:"createdBy"`
	CreatedOn  string      `json:"createdOn"`
	ModifiedBy string      `json:"modifiedBy"`
	ModifiedOn string      `json:"modifiedOn"`
}

type Parameter struct {
	ParameterID int    `json:"parameterID"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	CreatedBy   string `json:"createdBy"`
	CreatedOn   string `json:"createdOn"`
	ModifiedBy  string `json:"modifiedBy"`
	ModifiedOn  string `json:"modifiedOn"`
}

func (handler *ActivityHandler) GetAllActivitiesHandler(w http.ResponseWriter, r *http.Request) {
	a1 := Activity{
		ActivityID: 1,
		Name:       "Activity1",
		PrettyName: "Activity 1",
		Index:      0,
		Note:       "This is a note",
		Summary:    "Summary should go here",
		Hits:       0,
	}

	a2 := Activity{
		ActivityID: 2,
		Name:       "Activity2",
		PrettyName: "Activity 2",
		Index:      1,
		Note:       "This is a note for the other activity",
		Summary:    "Summary should go here and be really cool",
		Hits:       0,
	}

	helpers.RespondWithJSON(w, http.StatusOK, Activities{Activities: []Activity{a1, a2}})
}

func (handler *ActivityHandler) GetActivityByIDHandler(w http.ResponseWriter, r *http.Request) {
	p1 := Parameter{
		ParameterID: 1,
		Name:        "Columns",
		Value:       "A",
	}

	p2 := Parameter{
		ParameterID: 2,
		Name:        "Prefix",
		Value:       "-okay",
	}

	activity := Activity{
		ActivityID: 1,
		Name:       "AddPrefix",
		PrettyName: "Add Prefix",
		Index:      0,
		Note:       "This is a note",
		Summary:    "add -okay to column A",
		Hits:       0,
		Parameters: []Parameter{p1, p2},
	}

	helpers.RespondWithJSON(w, http.StatusOK, activity)
}
