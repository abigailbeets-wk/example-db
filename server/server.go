package server

import (
	"example-db/router"
	"fmt"
	"log"
	"net/http"
	"os"
)

// TODO maybe move this to main.go for DB injection

func StartServer() {
	port := "8080"
	r := router.NewRouter()

	if value, ok := os.LookupEnv("PORT"); ok {
		port = value
	}

	log.Println(fmt.Sprintf("Starting server on port %v...", port))
	log.Fatal(http.ListenAndServe(":"+port, r))
}
