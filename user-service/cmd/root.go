package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "user-service",
	Short: "user service application",
}

func Execute() error {
	registerCommands()

	return rootCmd.Execute()
}

func registerCommands() {
	rootCmd.AddCommand(
		serveCmd,
		migrateCmd,
		seedCmd)

	migrateCmd.AddCommand(
		migrateUpCmd,
		migrateDownCmd,
		migrateForceCmd,
		migrateVersionCmd)
}
