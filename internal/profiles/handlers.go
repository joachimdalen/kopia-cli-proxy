package profiles

import (
	"encoding/json"
	"net/http"

	"github.com/joachimdalen/kopia-cli-proxy/internal/kopia"
	"github.com/joachimdalen/kopia-cli-proxy/models"
)

func HandleListProfiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := kopia.GetProfiles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	j, err := json.Marshal(profiles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(j)
}

func HandleAddProfile(w http.ResponseWriter, r *http.Request) {
	req := &models.AddProfileRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	profile, err := kopia.AddProfile(*req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	j, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(j)
}
