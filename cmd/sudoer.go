/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// sudoerCmd represents the add-sudoer command
var sudoerCmd = &cobra.Command{
	Use:   "sudoer [username]",
	Short: "Add a user to sudoers (admins group)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		username := args[0]

		// Execute: sudo usermod -aG admins <username>
		command := exec.Command("sudo", "usermod", "-aG", "admins", username)

		// Capture output (optional but useful for debugging)
		output, err := command.CombinedOutput()
		if err != nil {
			fmt.Println("Error adding user to sudoers:", err)
			fmt.Println("Details:", string(output))
			return
		}

		fmt.Println("User successfully added to sudoers (admins group):", username)
	},
}

func init() {
	rootCmd.AddCommand(sudoerCmd)
}
