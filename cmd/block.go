package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var blockCmd = &cobra.Command{
	Use:   "block <username>",
	Short: "Block a user account",
	Long:  `Locks a user account using usermod -L, preventing them from logging in.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]

		out, err := exec.Command("sudo", "usermod", "-L", username).CombinedOutput()
		if err != nil {
			fmt.Printf("Error blocking user '%s': %v\n%s\n", username, err, string(out))
			return
		}

		fmt.Printf("User '%s' has been blocked successfully.\n", username)
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)
}
