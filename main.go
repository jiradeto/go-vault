package main

import (
	"fmt"
	"os"
	"vault/cobra"
)

func main() {
	if err := cobra.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
