package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func _responseWithErrorInfo(response http.ResponseWriter, err error, statusCode int) {
	errBuffer := []byte(err.Error())
	response.Header().Set("Content-Type", "text/plain")
	response.Header().Set("Content-Length", fmt.Sprint(len(errBuffer)))
	response.WriteHeader(statusCode)
	response.Write(errBuffer)
}

func _responseWithJSONData(response http.ResponseWriter, data []byte, statusCode int) {
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	response.Header().Set("Content-Length", fmt.Sprint(len(data)))
	response.WriteHeader(statusCode)
	response.Write(data)
}

func responseBadRequest(response http.ResponseWriter, err error) {
	_responseWithErrorInfo(response, err, http.StatusBadRequest)
}

// Response internal server error status code to the client
func responseInternalServerError(response http.ResponseWriter, err error) {
	_responseWithErrorInfo(response, err, http.StatusInternalServerError)
}

// Response OK status code
func responseOK(response http.ResponseWriter, data interface{}) {
	dataBuffer, err := loadDataAsBuffer(data)
	if err != nil {
		responseInternalServerError(response, err)
		return
	}

	_responseWithJSONData(response, dataBuffer, http.StatusOK)
}

// Response created status code and data created to the client
func responseCreated(response http.ResponseWriter, data interface{}) {
	dataBuffer, err := loadDataAsBuffer(data)
	if err != nil {
		responseInternalServerError(response, err)
		return
	}

	_responseWithJSONData(response, dataBuffer, http.StatusCreated)
}

func loadDataAsBuffer(data interface{}) (buffer []byte, err error) {
	if !isDataBuffer(data) {
		buffer, err = json.Marshal(data)
	} else {
		buffer = data.([]byte)
	}
	return
}

func isDataBuffer(data interface{}) bool {
	return fmt.Sprintf("%T", data) == "[]uint8"
}
