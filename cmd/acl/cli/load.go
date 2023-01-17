package cli

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	aclmgr "github.com/streamingfast/hivemapper_hdc_acl"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "load acl from disk",
	RunE:  loadRunE,
	Args:  cobra.ExactArgs(1),
}

func init() {
	RootCmd.AddCommand(loadCmd)
}

type loadResult struct {
	Acl       *aclmgr.Acl `json:"acl"`
	Signature []byte      `json:"signature"`
}

func loadRunE(cmd *cobra.Command, args []string) error {
	aclFolder := args[0]
	result := loadResult{}

	if aclmgr.AclExistOnDevice(aclFolder) {
		acl, signature, err := aclmgr.NewAclFromFile(aclFolder)
		if err != nil {
			panic(err)
		}
		result.Acl = acl
		result.Signature = signature[:]
	}

	data, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("unable to marshal result: %w", err)
	}
	hexData := hex.EncodeToString(data)
	fmt.Print(hexData) //this will be read as command output ...

	return nil
}
