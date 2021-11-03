package server

import (
	"encoding/base64"
	"net/http"
	"os"
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

func isRequestAuthorized(request *http.Request) bool {
	authEncoded := request.Header.Get("Authorization")
	if len(authEncoded) == 0 || strings.Index("Basic ", string(authEncoded)) == 0 {
		return false
	}

	authDecoded, err := base64.URLEncoding.DecodeString(authEncoded[6:])
	if err != nil {
		return false
	}

	auth := strings.Split(string(authDecoded), ":")
	if len(auth) != 2 {
		return false
	}

	if auth[0] != os.Getenv("AUTH_USER") || auth[1] != os.Getenv("AUTH_PASS") {
		return false
	}

	return true
}
