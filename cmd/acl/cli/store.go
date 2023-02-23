package cli

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
	aclmgr "github.com/streamingfast/hivemapper_hdc_acl"
	"github.com/streamingfast/solana-go"
)

var storeCmd = &cobra.Command{
	Use:   "store {hex_acl} {signature} {destination_path}",
	Short: "store acl to disk",
	RunE:  storeRunE,
	Args:  cobra.ExactArgs(3),
}

func init() {
	RootCmd.AddCommand(storeCmd)
}

func storeRunE(cmd *cobra.Command, args []string) error {
	aclArg := args[0]
	signArg := args[1]
	aclFolder := args[2]

	aclData, err := hex.DecodeString(aclArg)
	if err != nil {
		return fmt.Errorf("unable to decode acl: %w", err)
	}

	acl, err := aclmgr.NewAclFromData(aclData)
	if err != nil {
		return fmt.Errorf("unable to create acl: %w", err)
	}

	if !isGranted(acl, aclFolder) {
		return fmt.Errorf("acl access not granted")
	}

	signature, err := solana.NewSignatureFromBase58(signArg)
	if err != nil {
		return fmt.Errorf("unable to decode signature: %w", err)
	}

	if acl.ValidateStoreSignature(signature) {
		err = acl.Store(aclFolder, signature)
		if err != nil {
			return fmt.Errorf("unable to store acl: %w", err)
		}
		return nil
	}

	return fmt.Errorf("invalid signature")
}

func isGranted(newAcl *aclmgr.Acl, sourcePath string) bool {
	if !aclmgr.AclExistOnDevice(sourcePath) {
		return true
	}

	currentAcl, _, err := aclmgr.NewAclFromFile(sourcePath)
	if err != nil {
		panic(err)
	}

	for _, manager := range currentAcl.Managers {
		for _, newManager := range newAcl.Managers {
			if manager == newManager {
				return true
			}
		}
	}
	return false
}
