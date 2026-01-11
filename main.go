package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joachimdalen/kopia-cli-proxy/internal/environment"
	"github.com/joachimdalen/kopia-cli-proxy/internal/maint"
	"github.com/joachimdalen/kopia-cli-proxy/internal/profiles"
)

func main() {
	r := mux.NewRouter()

	// Profiles
	r.HandleFunc("/profiles", basicAuth(profiles.HandleListProfiles)).Methods("GET")
	r.HandleFunc("/profiles", basicAuth(profiles.HandleAddProfile)).Methods("POST")

	// Maintenance
	r.HandleFunc("/maintenance/info", basicAuth(maint.HandleGetMaintenanceInfo)).Methods("GET")

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	fmt.Println("Server starting on port 8080...")
	log.Fatal(srv.ListenAndServe())
}

func basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		username, err := environment.GetUsername()
		password, err := environment.GetPassword()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if !ok || user != username || pass != password {
			w.Header().Set("WWW-Authenticate", `Basic realm="kopia"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
