// handlers/handlers.go
package handlers

import (
	"net/http"
	"github.com/isaka-james/svg-go/models"
	"github.com/isaka-james/svg-go/utils"
	"log"
)

func SVGHandler(w http.ResponseWriter, r *http.Request) {
	// Validate request method
	if r.Method != http.MethodGet {
		utils.RespondJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	// Parse query parameters
	name1 := r.URL.Query().Get("name1")
	name2 := r.URL.Query().Get("name2")

	// Create SVG response
	svgResponse := models.SVGResponse{
		Name1: name1,
		Name2: name2,
	}

	// Generate SVG content
	svgContent := svgResponse.GenerateSVG()

	// Set Content-Type header to indicate SVG file
	w.Header().Set("Content-Type", "image/svg+xml")

	// Write SVG content to response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(svgContent))

	log.Printf("Generated SVG for names: %s and %s", name1, name2)
}
