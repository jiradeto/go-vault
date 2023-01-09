package cobra

import (
	"fmt"
	"log"
	"vault/secret"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("must provide one key name")
		}
		v := secret.NewVault(vaultKey, FILEPATH)
		val, err := v.Get(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Key:", args[0])
		fmt.Println("Value:", val)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
