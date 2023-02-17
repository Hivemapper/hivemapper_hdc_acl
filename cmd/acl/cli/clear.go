package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	aclmgr "github.com/streamingfast/hivemapper_hdc_acl"
)

var clearCmd = &cobra.Command{
	Use:   "clear {acl_path}",
	Short: "remove acl from disk",
	RunE:  clearRunE,
	Args:  cobra.ExactArgs(1),
}

func init() {
	RootCmd.AddCommand(clearCmd)
}

func clearRunE(cmd *cobra.Command, args []string) error {
	aclFolder := args[0]

	//todo: add signature parameter
	//todo: validate signature signature against acl

	if aclmgr.AclExistOnDevice(aclFolder) {
		if err := aclmgr.AclClearFromDevice(aclFolder); err != nil {
			return fmt.Errorf("unable to clear acl: %w", err)
		}
	}

	return nil
}
