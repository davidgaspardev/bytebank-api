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
					// Returns error from reading body data
					response.Header().Set("Content-Type", "text/plain; charset=utf-8")
					response.WriteHeader(http.StatusBadRequest)
					response.Write([]byte(err.Error()))
					return
				}

				data.DateTime = time.Now()

				if err := service.AddTransfer(&data); err != nil {
					response.Header().Set("Content-Type", "text/plain; charset=utf-8")
					response.WriteHeader(http.StatusBadRequest)
					response.Write([]byte(err.Error()))
					return
				}

				var resData, _ = json.Marshal(data)
				response.Header().Set("Content-Type", "application/json; charset=utf-8")
				response.WriteHeader(http.StatusOK)
				response.Write(resData)
			} else if request.Method == http.MethodGet {
				var data = service.GetAllTransfers()
				var body, err = json.Marshal(data)
				if err != nil {
					// There was a problem converting data to a byte array
					response.Header().Set("Content-Type", "text/plain; charset=utf-8")
					response.WriteHeader(http.StatusBadRequest)
					response.Write([]byte(err.Error()))
					break
				}
				response.Header().Set("Content-Type", "application/json; charset=utf-8")
				response.Header().Set("Content-Length", fmt.Sprint(len(body)))
				response.WriteHeader(http.StatusOK)
				response.Write(body)
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
