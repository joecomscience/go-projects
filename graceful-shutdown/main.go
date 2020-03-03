package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	 r := mux.NewRouter()

	 r.HandleFunc("/readiness", func(writer http.ResponseWriter, request *http.Request) {
		 writer.WriteHeader(http.StatusOK)
		 fmt.Fprintf(writer, "ok")
	 })
	 r.HandleFunc("/liveness", func(writer http.ResponseWriter, request *http.Request) {
		 writer.WriteHeader(http.StatusOK)
		 fmt.Fprint(writer, "ok")
	 })
	 r.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
		 fmt.Fprintf(writer, "hello v1")
	 })

	log.Fatal(http.ListenAndServe(":8080", r))
}