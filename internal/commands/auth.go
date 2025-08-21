package commands

import (
	"bufio"
	auth "dflow/internal/cloud/auth"
	"dflow/internal/persistency/repository"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Auth = &cobra.Command{
	Use:   "auth",
	Short: "Connects cli to cloud",
	Run: func(cmd *cobra.Command, args []string) {
		auth.LoginWeb()
		fmt.Println("You successfully logged in by your browser.\nPlease insert your customer code: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		repository.UpdateClientID(strings.TrimSpace(input))
	},
}

var AuthReset = &cobra.Command{
	Use:   "reset",
	Short: "Resets authorizations",
	Run: func(cmd *cobra.Command, args []string) {
		auth.Reset()
	},
}
