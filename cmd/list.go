/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List system users with UID >= 1000",
	Run: func(cmd *cobra.Command, args []string) {

		// Open the /etc/passwd file
		file, err := os.Open("/etc/passwd")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Create a scanner to read the file line by line
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()

			// Split the line by ":" to extract fields
			fields := strings.Split(line, ":")

			// Ensure the line has enough fields
			if len(fields) < 3 {
				continue
			}

			// Extract username and UID
			username := fields[0]
			uidStr := fields[2]

			// Convert UID from string to integer
			uid, err := strconv.Atoi(uidStr)
			if err != nil {
				continue
			}

			if username == "nobody" {
				continue
			}

			// Filter users with UID >= 1000
			if uid >= 1000 {
				fmt.Println(username)
			}

			
		}

		// Check for scanning errors
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}