package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var printCmd = &cobra.Command{
	Use: "print",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Configuration parameters:\n\t"+
			"Config file: %s\n\t"+
			"Api version: %s\n\t"+
			"Client ID: %s\n\t"+
			"Client Secret: %s\n\t"+
			"User: %s\n\t"+
			"Password: %s\n\t"+
			"Token: %s\n\t"+
			"Environment: %s\n",
			viper.GetString(ConfigKey),
			viper.GetString(ApiVersionKey),
			viper.GetString(ClientIdKey),
			viper.GetString(ClientSecretKey),
			viper.GetString(UserKey),
			viper.GetString(PasswordKey),
			viper.GetString(SecurityTokenKey),
			viper.GetString(EnvironmentKey))
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
