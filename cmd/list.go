package cmd

import (
	"github.com/bahelms/noted/config"
	"github.com/bahelms/noted/core"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all tracked files",
	Long:  "Long Description",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.New()
		core.ListFiles(cfg)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
