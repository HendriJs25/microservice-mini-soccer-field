package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "user-service",
	Short: "user service application",
}

func Execute() error {
	return rootCmd.Execute()
}
