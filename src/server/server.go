package server

import (
	"fmt"
	"net/http"
)

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()

	buildRoutes()
}

func Start(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
