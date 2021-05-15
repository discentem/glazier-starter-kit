package cmd

import (
	"errors"

	"github.com/discentem/glazier-config/cmd/sync"
	"github.com/discentem/glazier-config/cmd/unattend"
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
		return sync.Execute(args[0])
	},
}

var unattendCmd = &cobra.Command{
	Use:   "unattend",
	Short: "Generate Unattend with random password",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("unattend must have 2 args: [source] [destination]")
		}
		return unattend.Execute(args[0], args[1])
	},
}

var rootCmd = &cobra.Command{
	Use:   "things",
	Short: "tools for managing glazier repo",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(unattendCmd)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
