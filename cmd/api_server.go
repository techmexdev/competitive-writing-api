package main

import (
	"log"
	"net/http"
	"os"

	mockSrv "github.com/techmexdev/competitive_writing_api/pkg/server/mock"
	mockStore "github.com/techmexdev/competitive_writing_api/pkg/storage/mock"
)

func main() {
	mockEnvVar := os.Getenv("MOCK")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if mockEnvVar == "TRUE" {
		store := mockStore.New()
		server := mockSrv.New(store)
		log.Printf("Listening on port " + port)
		http.ListenAndServe(":"+port, server)
	}
}
