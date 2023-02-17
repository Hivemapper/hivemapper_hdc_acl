package cli

import (
	"fmt"

	"github.com/streamingfast/solana-go"

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

	if aclmgr.AclExistOnDevice(aclFolder) {
		acl, _, err := aclmgr.NewAclFromFile(aclFolder)
		if err != nil {
			return fmt.Errorf("unable to read acl: %w", err)
		}

		if acl.Version != "" && signatureB58 == "" {
			return fmt.Errorf("ACL on device requires a signature to be cleared")
		}

		if signatureB58 != "" {
			signature, err := solana.NewSignatureFromBase58(signatureB58)
			if err != nil {
				return fmt.Errorf("unable to decode signature: %w", err)
			}
			if !acl.ValidateSignature(signature) {
				return fmt.Errorf("invalid signature")
			}
		}

		if err := aclmgr.AclClearFromDevice(aclFolder); err != nil {
			return fmt.Errorf("unable to clear acl: %w", err)
		}
	}

	return nil
}
