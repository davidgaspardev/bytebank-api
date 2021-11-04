package server

import (
	model "bytebank-api/src/models"
	service "bytebank-api/src/service"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func logRequest(request *http.Request) {
	fmt.Printf("\n[ %s ] Path: %s\n", request.RemoteAddr, request.URL.Path)
	if len(request.URL.RawQuery) > 0 {
		fmt.Printf("[ %s ] Query: %s\n", request.RemoteAddr, request.URL.RawQuery)
	}
	if request.Header.Get("Content-Type") != "" {
		fmt.Printf("[ %s ] Content type: %s\n", request.RemoteAddr, request.Header.Get("Content-Type"))
	}
	if request.Header.Get("Accept") != "" {
		fmt.Printf("[ %s ] Accept: %s\n", request.RemoteAddr, request.Header.Get("Accept"))
	}
	if request.Header.Get("User-Agent") != "" {
		fmt.Printf("[ %s ] User agent: %s\n", request.RemoteAddr, request.Header.Get("User-Agent"))
	}
	if request.Header.Get("Authorization") != "" {
		fmt.Printf("[ %s ] Authorization: %s\n", request.RemoteAddr, request.Header.Get("Authorization"))
	}
}

func buildRoutes() {
	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

		logRequest(request)

		switch request.URL.Path {
		// Adding Cardflix in database
		case "/transactions":
			if request.Method == http.MethodPost && isRequestAuthorized(request) {
				var data model.Transfer
				// Get the requisition body data
				var err = json.NewDecoder(request.Body).Decode(&data)
				if err != nil {
					responseBadRequest(response, err)
					return
				}

				data.DateTime = time.Now()
				if err := service.AddTransfer(&data); err != nil {
					responseBadRequest(response, err)
					return
				}

				responseCreated(response, data)

			} else if request.Method == http.MethodGet {
				var data = service.GetAllTransfers()
				responseOK(response, data)
			} else {
				http.NotFound(response, request)
			}
			break
		default:
			http.NotFound(response, request)
			break
		}
	})
}
