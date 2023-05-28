package cmd

import (
	"github.com/bahelms/noted/config"
	"github.com/bahelms/noted/core"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Replace local files with files stored externally",
	Long:  "External storage is the source of truth. When running `sync`, all external files will be pulled down locally, creating new files or overwriting existing ones.",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Config{LocalStorageDir: ".noted", AwsProfile: "default", S3BucketName: "noted-file-storage"}
		core.SyncFiles(cfg)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
