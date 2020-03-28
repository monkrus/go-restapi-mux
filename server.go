package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, request *http.Request) {
		fmt.Println(resp, "up and running")
	})
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
