package maint

import (
	"encoding/json"
	"net/http"

	"github.com/joachimdalen/kopia-cli-proxy/internal/kopia"
)

func HandleListUsers(w http.ResponseWriter, r *http.Request) {
	maintInfo, err := kopia.GetMaintenanceInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	j, err := json.Marshal(maintInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(j)
}
