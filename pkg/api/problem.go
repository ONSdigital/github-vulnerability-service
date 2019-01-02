package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// Problem is as specified in RFC7807 - https://tools.ietf.org/html/rfc7807
type Problem struct {
	Type   string `json:"type,omitempty"`   // Link to a resource for the problem
	Title  string `json:"title,omitempty"`  // Short description of the issue
	Status int    `json:"status"`           // The http status code
	Detail string `json:"detail,omitempty"` // Further human-readable detail
}

// WriteProblemResponse writes an API problem response report to the given ResponseWriter.
// If for some reason it fails to marshal the json response, it returns a 500
// internal error.
func WriteProblemResponse(problem Problem, rw http.ResponseWriter) {
	pr, err := json.Marshal(&problem)
	if err != nil {
		log.Printf(`event="Error writing problem reponse" error="%v"`, err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if problem.Status == 0 {
		problem.Status = http.StatusInternalServerError
	}
	rw.Header().Set("Content-Type", "application/problem+json")
	rw.Header().Set("Content-Language", "en")
	rw.WriteHeader(problem.Status)
	rw.Write(pr)
	return
}
