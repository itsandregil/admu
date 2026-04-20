package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

const defaultPassword = "Culantro123!"

var addCmd = &cobra.Command{
	Use:   "add <username>",
	Short: "Add a new user to your machine.",
	Long:  `Creates a new user with home directory, sets a default password, and configures password expiration policy.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]

		// 1. useradd -m <USERNAME>
		fmt.Printf("Creating user '%s'...\n", username)
		out, err := exec.Command("sudo", "useradd", "-m", username).CombinedOutput()
		if err != nil {
			fmt.Printf("Error creating user: %v\n%s\n", err, string(out))
			return
		}
		fmt.Println("User created successfully.")

		// 2. echo "<USERNAME>:<DEFAULT_PASS>" | chpasswd
		fmt.Println("Setting default password...")
		chpasswd := exec.Command("sudo", "chpasswd")
		chpasswd.Stdin = strings.NewReader(fmt.Sprintf("%s:%s", username, defaultPassword))
		out, err = chpasswd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error setting password: %v\n%s\n", err, string(out))
			return
		}
		fmt.Println("Password set successfully.")

		// 3. chage -d 0 -W 10 -M 30 -I 0 <USERNAME>
		fmt.Println("Configuring password expiration policy...")
		out, err = exec.Command("sudo", "chage", "-d", "0", "-W", "10", "-M", "30", "-I", "0", username).CombinedOutput()
		if err != nil {
			fmt.Printf("Error configuring password policy: %v\n%s\n", err, string(out))
			return
		}
		fmt.Println("Password policy configured successfully.")

		fmt.Printf("\nUser '%s' created and configured. They will be prompted to change their password on first login.\n", username)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
