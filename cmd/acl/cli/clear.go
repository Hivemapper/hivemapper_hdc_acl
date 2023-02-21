package cli

import (
	"github.com/spf13/cobra"
	aclmgr "github.com/streamingfast/hivemapper_hdc_acl"
)

var clearCmd = &cobra.Command{
	Use:   "clear {acl_path} {signature}",
	Short: "remove acl from disk",
	RunE:  clearRunE,
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	RootCmd.AddCommand(clearCmd)
}

func clearRunE(cmd *cobra.Command, args []string) error {
	aclFolder := args[0]
	signatureB58 := ""
	if len(args) > 1 {
		signatureB58 = args[1]
	}

	err := aclmgr.AclClearFromDevice(aclFolder, signatureB58)
	if err != nil {
		return err
	}

	return nil
}
