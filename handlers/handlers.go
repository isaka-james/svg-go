// handlers/handlers.go
package handlers

import (
	"net/http"
	"github.com/isaka-james/svg-go/models"
	"github.com/isaka-james/svg-go/utils"
	"log"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	// Validate request method
	if r.Method != http.MethodPost {
		utils.RespondJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	// Parse JSON request into AddRequest model
	var request models.AddRequest
	err := utils.ParseJSONRequest(r, &request)
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	// Validate input data
	if request.Number1 < 0 || request.Number2 < 0 {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "Input numbers must be positive"})
		return
	}

	// Perform addition
	result := request.Number1 + request.Number2

	// Create response
	response := models.AddResponse{Result: result}

	// Send JSON response
	utils.RespondJSON(w, http.StatusOK, response)
	log.Printf("Added %d and %d, result: %d", request.Number1, request.Number2, result)
}
