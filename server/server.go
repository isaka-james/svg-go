// server/server.go
package server

import (
	"log"
	"net/http"

	"github.com/isaka-james/svg-go/handlers"
)

func StartServer() error {
	http.HandleFunc("/api/add", handlers.AddHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return err
	}
	return nil
}
