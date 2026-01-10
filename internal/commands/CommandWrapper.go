package commands

import (
	"log"
	"os/exec"
)

func ExecuteAndGetResponse(command string) ([]byte, error) {
	// Define the command and its arguments separately
	cmd := exec.Command("powershell.exe", command)

	// Run the command and wait for completion
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd failed with %s\n", err)
		return nil, err
	}

	jsonDataBytes := []byte(out)
	return jsonDataBytes, nil
}
