package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func battleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Battle!\n"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/battles", battleHandler)

	log.Fatal(http.ListenAndServe(":13499", r))
}
