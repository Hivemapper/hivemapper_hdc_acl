package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "return the acl manager version",
	RunE:  versionRunE,
	Args:  cobra.ExactArgs(0),
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

func versionRunE(cmd *cobra.Command, args []string) error {
	fmt.Print("1.1.2")
	return nil
}
