package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/melan/go-force/force"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	DefaultSfdcApi   = "v43.0"
	DefaultSfdcEnv   = "prod"
	ConfigKey        = "config"
	ApiVersionKey    = "api-version"
	ClientIdKey      = "client-id"
	ClientSecretKey  = "client-secret"
	UserKey          = "username"
	PasswordKey      = "password"
	SecurityTokenKey = "security-token"
	EnvironmentKey   = "environment"
	TraceOnKey       = "trace-on"
)

var (
	rootCmd = &cobra.Command{}

	cfgFile                string
	apiVersion             string
	clientId, clientSecret string
	username, password     string
	securityToken          string
	environment            string
	traceOn                bool
)

type screenLogger struct{}

func (cl *screenLogger) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, ConfigKey, "./config.yaml", "config file")
	rootCmd.PersistentFlags().StringVar(&apiVersion, ApiVersionKey, DefaultSfdcApi, "Api version for Salesforce, must be in a format vXX.X")
	rootCmd.PersistentFlags().StringVar(&clientId, ClientIdKey, "", "Client ID from a Connected App to connect to Salesforce API")
	rootCmd.PersistentFlags().StringVar(&clientSecret, ClientSecretKey, "", "Client Secret form a Connected app to connect to Salesforce")
	rootCmd.PersistentFlags().StringVar(&username, UserKey, "", "Username to connect to Salesforce")
	rootCmd.PersistentFlags().StringVar(&password, PasswordKey, "", "Password to connect to Salesforce")
	rootCmd.PersistentFlags().StringVar(&securityToken, SecurityTokenKey, "", "Security Token to connect to Salesforce")
	rootCmd.PersistentFlags().StringVar(&environment, EnvironmentKey, DefaultSfdcEnv, "Environment to connect to Salesforce")
	rootCmd.PersistentFlags().BoolVar(&traceOn, TraceOnKey, false, "Enable tracing of all HTTP communication")

	viper.SetEnvPrefix("SFDC")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	viper.BindEnv(ApiVersionKey)
	viper.BindEnv(ClientIdKey)
	viper.BindEnv(ClientSecretKey)
	viper.BindEnv(UserKey)
	viper.BindEnv(PasswordKey)
	viper.BindEnv(SecurityTokenKey)
	viper.BindEnv(EnvironmentKey)
	viper.BindEnv(TraceOnKey)

	viper.BindPFlag(ApiVersionKey, rootCmd.PersistentFlags().Lookup(ApiVersionKey))
	viper.BindPFlag(ClientIdKey, rootCmd.PersistentFlags().Lookup(ClientIdKey))
	viper.BindPFlag(ClientSecretKey, rootCmd.PersistentFlags().Lookup(ClientSecretKey))
	viper.BindPFlag(UserKey, rootCmd.PersistentFlags().Lookup(UserKey))
	viper.BindPFlag(PasswordKey, rootCmd.PersistentFlags().Lookup(PasswordKey))
	viper.BindPFlag(SecurityTokenKey, rootCmd.PersistentFlags().Lookup(SecurityTokenKey))
	viper.BindPFlag(EnvironmentKey, rootCmd.PersistentFlags().Lookup(EnvironmentKey))
	viper.BindPFlag(TraceOnKey, rootCmd.PersistentFlags().Lookup(TraceOnKey))
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("/etc/forcecom")
		viper.AddConfigPath("$HOME/.forcecom")
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
	}
}

func initClient() (*force.Api, error) {
	oauth, err := force.CreateOAuth(viper.GetString(ClientIdKey), viper.GetString(ClientSecretKey),
		viper.GetString(UserKey), viper.GetString(PasswordKey), viper.GetString(SecurityTokenKey),
		viper.GetString(EnvironmentKey))
	if err != nil {
		return nil, fmt.Errorf("can't construct OAuth object: %v", err)
	}

	api, err := force.CreateApi(apiVersion, oauth)
	if err != nil {
		return nil, fmt.Errorf("can't construct API Object: %v", err)
	}

	if traceOn {
		logger := &screenLogger{}
		api.TraceOn("", logger)
	}

	return api, nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
