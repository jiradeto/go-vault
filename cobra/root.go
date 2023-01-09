package cobra

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Secret is an API key and other secrets manager",
}

const FILEPATH = ".secret.key"

var (
	vaultKey string
)

func init() {
	RootCmd.PersistentFlags().StringVar(&vaultKey, "key", "", "key for secrets encryption/decryption")
	RootCmd.MarkPersistentFlagRequired("key")
}
