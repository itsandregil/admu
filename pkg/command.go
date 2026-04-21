package admu

import (
	"fmt"
	"os/exec"
)

func RunCommand(cmd string, args ...string) {
	arguments := getArguments(cmd, args)
	command := exec.Command("sudo", arguments...)
	if err := command.Run(); err != nil {
		fmt.Printf("command failed: %v\n", err)
	}
}

func getArguments(cmd string, args []string) []string {
	arguments := []string{}
	arguments = append(arguments, cmd)
	arguments = append(arguments, args...)
	return arguments
}
