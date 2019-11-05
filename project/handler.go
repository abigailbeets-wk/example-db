package project

import (
	"example-db/helpers"
	"log"
	"net/http"
	"os"
)

type ProjectHandler struct {
	Service IProjectService
}

type Projects struct {
	Projects []Project `json:"projects"`
}

func (handler *ProjectHandler) GetAllProjectsHandler(w http.ResponseWriter, r *http.Request) {
	// supposing this was part of the request body, we'd add a struct tag to the
	// request body struct containing a constraint on the accountID and ensure the
	// accountID is safe to use in our query and reject if
	// it does not meet requirements around length or another criteria
	// such as being only alphanumeric characters. For this example, just
	// going to check if it's 32 characters long and not make it part of
	// the request body.
	accountID := os.Getenv("ACCOUNT_ID")
	if len(accountID) < 32 {
		helpers.RespondWithStatus(w, http.StatusForbidden)
		return
	}

	projects, err := handler.Service.GetAllProjects(accountID)
	if err != nil {
		log.Println(err)
		helpers.RespondWithError(w, http.StatusInternalServerError, "error retrieving projects")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, Projects{Projects: projects})
}
