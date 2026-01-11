package environment

import (
	"errors"
	"os"
)

func GetKopiaContainerName() (string, error) {
	containerName := os.Getenv("KCP_CONTAINER_NAME")
	if containerName == "" {
		return "", errors.New("Failed to find KCP_CONTAINER_NAME")
	}
	return containerName, nil
}

func GetUsername() (string, error) {
	username := os.Getenv("KCP_USERNAME")
	if username == "" {
		return "", errors.New("Failed to find KCP_USERNAME")
	}
	return username, nil
}

func GetPassword() (string, error) {
	password := os.Getenv("KCP_PASSWORD")
	if password == "" {
		return "", errors.New("Failed to find KCP_PASSWORD")
	}
	return password, nil
}
