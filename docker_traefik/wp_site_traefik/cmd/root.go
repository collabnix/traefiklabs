package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
}

// RootCmd is the root CLI command
var RootCmd = &cobra.Command{
	Use:           "pekka",
	Short:         "Deploy and manage multiple wordpress sites with traefik and docker",
	SilenceUsage:  true,
	SilenceErrors: true,
}
