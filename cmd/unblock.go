/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// unblockCmd represents the unblock command
var unblockCmd = &cobra.Command{
	Use:   "unblock <username>",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		out, err := exec.Command("sudo", "usermod", "-U", username).CombinedOutput()
		if err != nil {
			fmt.Printf("Error unblocking user '%s': %v\n%s\n", username, err, string(out))
			return
		}

		fmt.Printf("User '%s' has been unblocked successfully.\n", username)
	},
}

func init() {
	rootCmd.AddCommand(unblockCmd)
}
