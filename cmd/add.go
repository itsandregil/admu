package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	admu "github.com/itsandregil/admu/pkg"
	"github.com/spf13/cobra"
)

const defaultPassword = "Culantro123"

var addCmd = &cobra.Command{
	Use:   "add [username]",
	Short: "Add a new user to your machine.",
	Long:  `Creates a new user with home directory, sets a default password, and configures password expiration policy.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		admu.RunCommand("useradd", "-m", username)
		chpasswd := exec.Command("sudo", "chpasswd")
		chpasswd.Stdin = strings.NewReader(fmt.Sprintf("%s:%s", username, defaultPassword))
		out, err := chpasswd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error setting password: %v\n%s\n", err, string(out))
			return
		}
		admu.RunCommand("chage", "-d", "0", "-W", "10", "-M", "30", "-I", "0", username)
		fmt.Printf("User '%s' created and configured.\n", username)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
