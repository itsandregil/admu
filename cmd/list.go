package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users in the system",
	Long:  `Read users inside the /etc/passwd file, parse its output and shows all users in the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open("/etc/passwd")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Split(line, ":")
			if len(fields) < 3 {
				continue
			}
			username := fields[0]
			uid, err := strconv.Atoi(fields[2])
			if err != nil {
				continue
			}
			if uid >= 1000 && uid <= 60000 {
				fmt.Println(username)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
