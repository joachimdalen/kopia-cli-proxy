package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joachimdalen/kopia-cli-proxy/internal/maint"
	"github.com/joachimdalen/kopia-cli-proxy/internal/users"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", users.HandleListUsers).Methods("GET")
	r.HandleFunc("/maintenance/info", maint.HandleListUsers).Methods("GET")

	srv := &http.Server{
		Addr:    "localhost:1212",
		Handler: r,
	}

	fmt.Println("Server starting on port 1212...")
	log.Fatal(srv.ListenAndServe())
}
