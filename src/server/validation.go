package server

import (
	"net/http"
	"strings"
)

// Checking if request is application/json
func isRequestApplicationJSON(request *http.Request) bool {
	if request.Method != http.MethodPost {
		return false
	}
	if request.Header.Get("Content-Type") != "application/json" {
		return false
	}
	return true
}

// Checking if request is multipart/form-data
func isRequestMulpartFormData(request *http.Request) bool {
	if request.Method != http.MethodPost {
		return false
	}
	if strings.Index(request.Header.Get("Content-Type"), "multipart/form-data; boundary=") != 0 {
		return false
	}
	return true
}
