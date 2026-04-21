/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	admu "github.com/itsandregil/admu/pkg"
	"github.com/spf13/cobra"
)

var sudoerCmd = &cobra.Command{
	Use:   "sudoer [username]",
	Short: "Add a user to the sudoers group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		admu.RunCommand("usermod", "-aG", "admins", username)
		fmt.Println("User successfully added to sudoers:", username)
	},
}

func init() {
	rootCmd.AddCommand(sudoerCmd)
}
