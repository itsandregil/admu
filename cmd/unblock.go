package cmd

import (
	"fmt"

	admu "github.com/itsandregil/admu/pkg"
	"github.com/spf13/cobra"
)

var unblockCmd = &cobra.Command{
	Use:   "unblock [username]",
	Short: "Unblocks a user in the machine",
	Long:  `Unlocks a user account using usermod -U, allowing them to log in.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		admu.RunCommand("usermod", "-U", username)
		fmt.Printf("User '%s' has been unblocked.\n", username)
	},
}

func init() {
	rootCmd.AddCommand(unblockCmd)
}
