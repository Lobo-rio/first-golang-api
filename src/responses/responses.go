package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON function that returns a JSON response to the request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}

}

// Error function that returns an error in JSON format
func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Err string `json:"err"`
	}{
		Err: err.Error(),
	})
}