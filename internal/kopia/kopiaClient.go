package kopia

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joachimdalen/kopia-cli-proxy/internal/commands"
	"github.com/joachimdalen/kopia-cli-proxy/models"
)

func GetProfiles() (*models.ProfilesResponse, error) {
	var cliProfiles []*models.KopiaProfile
	containerName := os.Getenv("KCP_CONTAINER_NAME")
	bytes, err := commands.ExecuteAndGetResponse(fmt.Sprintf("docker exec %s kopia server users list --json", containerName))

	err = json.Unmarshal(bytes, &cliProfiles)
	if err != nil {
		log.Fatalf("error unmarshalling JSON: %v", err)
		return nil, err
	}

	if cliProfiles == nil {
		cliProfiles = []*models.KopiaProfile{}
	}

	resp := &models.ProfilesResponse{
		Profiles: []*models.Profile{},
	}

	for _, pol := range cliProfiles {
		var parts = strings.Split(pol.Username, "@")
		resp.Profiles = append(resp.Profiles, &models.Profile{
			Username: pol.Username,
			Hostname: parts[1],
			User:     parts[0],
		})
	}

	return resp, err
}
func GetMaintenanceInfo() (*models.MaintenanceInfo, error) {
	var info *models.MaintenanceInfo
	containerName := os.Getenv("KCP_CONTAINER_NAME")
	bytes, err := commands.ExecuteAndGetResponse(fmt.Sprintf("docker exec %s kopia maintenance info --json", containerName))

	err = json.Unmarshal(bytes, &info)
	if err != nil {
		log.Fatalf("error unmarshalling JSON: %v", err)
		return nil, err
	}

	return info, err
}
