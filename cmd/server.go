package cmd

import (
	"github.com/oidc-soma/aerosquirrel/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Long:  "Start the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		address, err := cmd.Flags().GetString("listen-address")
		if err != nil {
			return err
		}

		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			return err
		}

		err = server.Exec(address, port)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().StringP("listen-address", "a", "0.0.0.0", `Listen address to start server on.`)
	serverCmd.PersistentFlags().Int("port", 8080, `Port to start server on, default:"8080".`)
}
