package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// APIResponse is the standard response for all API responses

type APIResponse struct {
	Success  bool 		`json:"success"`
	Data	 interface{}`json:"data,omitempty"`
	Error 	 string     `json:"error,omitempty"`
}

// JSON writes responses with proper headers
func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)


	if err := json.NewEncoder(w).Encode(APIResponse{
		Success: status < 400,
		Data: data,
	}); err != nil {
		log.Printf("error encoding repsonse: %v", err)
	}
}


// Error writes standardized error response
func Error(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(APIResponse{
		Success: false,
		Error: message,
	}); err != nil {
		log.Printf("error encoding error response: %v", err)
	}
}

