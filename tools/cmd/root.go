package cmd

import (
	"errors"
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
	Use:   "",
	Short: "tools for managing glazier repo",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func defaultEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Execute() error {
	rootCmd.AddCommand(syncCmd)

	syncCmd.Flags().String("root", "./../glazier-repo/.", "Path to start recursive s3 sync from")
	accessKey := syncCmd.Flags().String("access_key", os.Getenv("ACCESS_KEY"), "AWS Access Key")
	if *accessKey == "" {
		return errors.New("access_key must not be an empty string")
	}
	secretKey := syncCmd.Flags().String("secret_key", os.Getenv("SECRET_KEY"), "AWS Secret Key")
	if *secretKey == "" {
		return errors.New("secret_key must not be an empty string")
	}
	bucket := syncCmd.Flags().String("bucket_name", os.Getenv("BUCKET_NAME"), "AWS Bucket Name")
	if *bucket == "" {
		return errors.New("bucket_name must not be an empty string")
	}
	syncCmd.Flags().String("region", defaultEnv("REGION", "us-east-1"), "AWS Bucket Name")
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
