// Package commands: Terminal UI commands
package commands

import (
	"github.com/spf13/cobra"
)

var ConfigGet = &cobra.Command{
	Use:   "get",
	Short: "Gets some infomation form the environment",
}
