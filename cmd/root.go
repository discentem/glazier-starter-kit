package cmd

import (
	"github.com/discentem/glazier-config/cmd/sync"
	"github.com/discentem/glazier-config/cmd/unattend"
	"github.com/sethvargo/go-password/password"
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
		src, err := cmd.LocalFlags().GetString("source")
		if err != nil {
			return err
		}
		dest, err := cmd.LocalFlags().GetString("destination")
		if err != nil {
			return err
		}
		homePage, err := cmd.LocalFlags().GetString("home_page")
		if err != nil {
			return err
		}
		oemSupport, err := cmd.LocalFlags().GetString("oem_support_url")
		if err != nil {
			return err
		}
		registeredCo, err := cmd.LocalFlags().GetString("registered_co")
		if err != nil {
			return err
		}
		pass, err := cmd.LocalFlags().GetString("password")
		if err != nil {
			return err
		}
		if pass == "" {
			pass, err = password.Generate(60, 20, 0, false, true)
			if err != nil {
				return err
			}
		}
		s := unattend.Settings{
			Source:       src,
			Destination:  dest,
			HomePage:     homePage,
			OEMSupport:   oemSupport,
			RegisteredCo: registeredCo,
			Pass:         pass,
		}

		return s.Execute()
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

	unattendCmd.Flags().String("source", "unattend.xml", "Source unattend.xml file")
	unattendCmd.Flags().String("destination", "generated_unattend.xml", "Generated unattend.xml")
	unattendCmd.Flags().String("home_page", "https://google.com", "Microsoft-Windows-IE-InternetExplorer Home_Page")
	unattendCmd.Flags().String("oem_support_url", "https://dell.com", "OEMInformation SupportURL")
	unattendCmd.Flags().String("registered_co", "ACME", "RegisteredOrganization")
	unattendCmd.Flags().String("password", "", "Administrator password. If no password is provided, will be randomized.")

	rootCmd.AddCommand(unattendCmd)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
