package cmd

import (
	"fmt"

	"github.com/discentem/glazier-config/cmd/sync"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync Glazier configs and resources to s3",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			args = make([]string, 1)
			args[0] = "."
		}
		sync.Execute(args[0])
		return nil
	},
}

var rootCmd = &cobra.Command{
	Use:   "things",
	Short: "tools for managing glazier repo",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("blarg")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
