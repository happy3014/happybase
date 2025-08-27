package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "happybase_cli",
	Short: "happybase_cli",
	Long:  "命令行工具",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	rootCmd.AddCommand(GetConfigMigrateCmd())

	rootCmd.Execute()
}

func main() {
	Execute()
}
