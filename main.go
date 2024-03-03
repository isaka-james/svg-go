// main.go
package main

import (
	"log"
	"github.com/isaka-james/svg-go/server"
)

func main() {
	err := server.StartServer()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
