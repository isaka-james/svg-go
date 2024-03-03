// server/server.go
package server

import (
	"log"
	"net/http"
	"os"
	"github.com/isaka-james/svg-go/handlers"
)

func StartServer() error {

	port := os.Getenv("PORT")
	if port == "" {
	        port = "8080"
	}
	
	http.HandleFunc("/api/add", handlers.AddHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return err
	}
	return nil
}
