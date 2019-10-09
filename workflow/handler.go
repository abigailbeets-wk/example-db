package workflow

import (
	"example-db/activity"
	"example-db/helpers"
	"net/http"
)

type WorkflowHandler struct{}

type Workflow struct {
	WorkflowID int                 `json:"workflowID"`
	Name       string              `json:"name"`
	Activities []activity.Activity `json:"activities,omitempty"`
}

func (handler *WorkflowHandler) GetAllWorkflowsHandler(w http.ResponseWriter, r *http.Request) {
	w1 := Workflow{
		WorkflowID: 1,
		Name:       "My Workflow",
	}

	w2 := Workflow{
		WorkflowID: 2,
		Name:       "My Other Workflow",
	}

	workflows := []Workflow{w1, w2}

	helpers.RespondWithJSON(w, http.StatusOK, workflows)
}

func (handler *WorkflowHandler) GetWorkflowByIDHandler(w http.ResponseWriter, r *http.Request) {
	a1 := activity.Activity{
		ActivityID: 1,
		Name:       "Activity1",
		PrettyName: "Activity 1",
		Index:      0,
		Note:       "This is a note",
		Summary:    "Summary should go here",
		Hits:       0,
	}

	a2 := activity.Activity{
		ActivityID: 2,
		Name:       "Activity2",
		PrettyName: "Activity 2",
		Index:      1,
		Note:       "This is a note for the other activity",
		Summary:    "Summary should go here and be really cool",
		Hits:       0,
	}

	workflow := Workflow{
		WorkflowID: 1,
		Name:       "My Workflow",
		Activities: []activity.Activity{a1, a2},
	}

	helpers.RespondWithJSON(w, http.StatusOK, workflow)
}
