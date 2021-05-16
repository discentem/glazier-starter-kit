package cmd

import (
	"os"

	"github.com/discentem/glazier-config/cmd/sync"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync Glazier configs and resources to s3",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := cmd.LocalFlags().GetString("bucket_name")
		if err != nil {
			return err
		}
		k, err := cmd.LocalFlags().GetString("access_key")
		if err != nil {
			return err
		}
		s, err := cmd.LocalFlags().GetString("secret_key")
		if err != nil {
			return err
		}
		region, err := cmd.LocalFlags().GetString("region")
		if err != nil {
			return err
		}
		r, err := cmd.LocalFlags().GetString("root")
		if err != nil {
			return err
		}
		return sync.Execute(b, k, s, region, r)
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

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.Flags().String("root", ".", "Path to start recursive s3 sync from")
	syncCmd.Flags().String("access_key", os.Getenv("ACCESS_KEY"), "AWS Access Key")
	syncCmd.Flags().String("secret_key", os.Getenv("SECRET_KEY"), "AWS Secret Key")
	syncCmd.Flags().String("bucket_name", os.Getenv("BUCKET_NAME"), "AWS Bucket Name")
	syncCmd.Flags().String("region", getEnv("REGION", "us-east-1"), "AWS Bucket Name")
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
