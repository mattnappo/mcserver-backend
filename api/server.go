package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func test(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "test worked!")
}

func InitializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/api/createServer", CreateServer).Methods("POST")
	router.HandleFunc("/api/serverStatus/{hash}", CreateServer).Methods("GET")
}
