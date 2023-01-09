package cobra

import (
	"log"
	"vault/secret"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatal("must provide key and value pair")
		}
		v := secret.NewVault(vaultKey, FILEPATH)
		err := v.Set(args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
