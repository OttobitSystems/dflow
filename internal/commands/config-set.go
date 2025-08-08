// Package commands: Terminal UI commands
package commands

import (
	"github.com/spf13/cobra"
)

var ConfigSet = &cobra.Command{
	Use:   "set",
	Short: "Sets some infomation into the environment",
}
