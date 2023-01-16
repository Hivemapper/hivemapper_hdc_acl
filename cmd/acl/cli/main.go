package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/streamingfast/dlauncher/flags"
)

var RootCmd = &cobra.Command{Use: "acl", Short: "hdc acl manager"}
var allFlags = make(map[string]bool) // used as global because of async access to cobra init functions

func Main() {
	cobra.OnInitialize(func() {
		allFlags = flags.AutoBind(RootCmd, "acl")
	})

	if err := RootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
