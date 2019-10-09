package router

import (
	"example-db/activity"
	"example-db/project"
	"example-db/workflow"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	initializeRoutes(r)

	return r
}

func initializeRoutes(r *chi.Mux) {
	projectHandler := project.ProjectHandler{}
	workflowHandler := workflow.WorkflowHandler{}
	activityHandler := activity.ActivityHandler{}

	r.Route("/v1", func(r chi.Router) {
		r.Route("/parsing-projects", func(r chi.Router) {
			r.Get("/", projectHandler.GetAllProjectsHandler)
			r.Get("/{projectID}", projectHandler.GetProjectByIDHandler)
			r.Post("/", projectHandler.CreateProjectHandler)

			r.Route("/workflows", func(r chi.Router) {
				r.Get("/", workflowHandler.GetAllWorkflowsHandler)
				r.Get("/{workflowID}", workflowHandler.GetWorkflowByIDHandler)

				r.Route("/activities", func(r chi.Router) {
					r.Get("/", activityHandler.GetAllActivitiesHandler)
					r.Get("/{activityID}", activityHandler.GetActivityByIDHandler)
				})
			})
		})
	})
}
