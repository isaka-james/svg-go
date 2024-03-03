// models/models.go
package models

// AddRequest represents the structure of the request for adding two numbers
type AddRequest struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

// AddResponse represents the structure of the response for adding two numbers
type AddResponse struct {
	Result int `json:"result"`
}
