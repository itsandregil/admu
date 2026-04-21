package cmd

import (
	"fmt"

	admu "github.com/itsandregil/admu/pkg"
	"github.com/spf13/cobra"
)

var blockCmd = &cobra.Command{
	Use:   "block [username]",
	Short: "Block a user account",
	Long:  `Locks a user account using usermod -L, preventing them from logging in.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		admu.RunCommand("usermod", "-L", username)
		fmt.Printf("User '%s' has been blocked successfully.\n", username)
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)
}
