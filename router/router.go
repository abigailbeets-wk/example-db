package router

import (
	"database/sql"
	"example-db/project"
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	initializeRoutes(r)

	return r
}

func initializeRoutes(r *chi.Mux) {
	db := initializeDatabase()

	projectRepository := project.ProjectRepository{DB: db}
	projectService := project.ProjectService{Repository: &projectRepository}
	projectHandler := project.ProjectHandler{Service: &projectService}

	r.Route("/v1", func(r chi.Router) {
		r.Route("/parsing-projects", func(r chi.Router) {
			r.Get("/", projectHandler.GetAllProjectsHandler)
		})
	})
}

func initializeDatabase() *sql.DB {
	username := os.Getenv("DATABASE_USERNAME")
	pass := os.Getenv("DATABASE_PASSWORD")
	dbPort := os.Getenv("DATABASE_PORT")
	dbHost := os.Getenv("DATABASE_HOST")
	dbName := os.Getenv("DATABASE_NAME")

	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, pass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	// TODO remember how I used to do this the right way?
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully pinged database")
	return db
}
