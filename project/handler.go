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
