package kopia

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/joachimdalen/kopia-cli-proxy/internal/commands"
	"github.com/joachimdalen/kopia-cli-proxy/internal/environment"
	"github.com/joachimdalen/kopia-cli-proxy/models"
)

func GetProfiles() (*models.ProfilesResponse, error) {
	var cliProfiles []*models.KopiaProfile
	containerName, err := environment.GetKopiaContainerName()
	if err != nil {
		return nil, err
	}

	bytes, err := commands.ExecuteAndGetResponse(fmt.Sprintf("docker exec %s kopia server users list --json", containerName))
	err = json.Unmarshal(bytes, &cliProfiles)
	if err != nil {
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

func GetProfile(username string) (*models.Profile, error) {
	var cliProfile *models.KopiaProfile
	containerName, err := environment.GetKopiaContainerName()
	if err != nil {
		return nil, err
	}

	bytes, err := commands.ExecuteAndGetResponse(fmt.Sprintf("docker exec %s kopia server users info %s", containerName, username))
	err = json.Unmarshal(bytes, &cliProfile)
	if err != nil {
		return nil, err
	}

	var parts = strings.Split(cliProfile.Username, "@")
	return &models.Profile{
		Username: cliProfile.Username,
		Hostname: parts[1],
		User:     parts[0],
	}, nil

}

func AddProfile(profile models.AddProfileRequest) (*models.Profile, error) {
	containerName, err := environment.GetKopiaContainerName()
	if err != nil {
		return nil, err
	}

	bytes, err := commands.ExecuteAndGetResponse(fmt.Sprintf("docker exec %s kopia server users add %s --user-password \"%s\"", containerName, profile.Username, profile.Password))
	str := string(bytes)

	if strings.Contains(str, "Updated user credentials will take effect") {
		return GetProfile(profile.Username)
	}
	return nil, fmt.Errorf("Failed to create user: %s", str)
}

func GetMaintenanceInfo() (*models.MaintenanceInfo, error) {
	var info *models.MaintenanceInfo
	containerName, err := environment.GetKopiaContainerName()
	if err != nil {
		return nil, err
	}

	bytes, err := commands.ExecuteAndGetResponse(fmt.Sprintf("docker exec %s kopia maintenance info --json", containerName))

	err = json.Unmarshal(bytes, &info)
	if err != nil {
		log.Fatalf("error unmarshalling JSON: %v", err)
		return nil, err
	}

	return info, err
}
