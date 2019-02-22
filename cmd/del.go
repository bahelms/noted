package cmd

import (
	"github.com/bahelms/noted/config"
	"github.com/bahelms/noted/core"
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Removes files locally and remotely",
	Long:  "Long Description",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cfg := config.Config{LocalStorageDir: ".noted"}
			core.DeleteFile(cfg, args[0])
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
