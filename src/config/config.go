package config

import (
	"github.com/spf13/viper"
)

var (
	AuthServiceURL            string
	ProjectManagerServiceURL  string
	AllowUnsignedCertificates bool
)

func init() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")

	// Do not handle errors here because it is not mandatory
	// to have a config file. It should work fine with default values.
	viper.ReadInConfig()

	setDefaultValues()
	configureVariables()
}

func setDefaultValues() {
	viper.SetDefault("auth.url", "")           // TODO fill after it is determined
	viper.SetDefault("projectManager.url", "") // TODO fill after it is determined
	viper.SetDefault("allowUnsignedCert", false)
}

func configureVariables() {
	AuthServiceURL = viper.GetString("auth.url")
	ProjectManagerServiceURL = viper.GetString("projectManager.url")
	AllowUnsignedCertificates = viper.GetBool("allowUnsignedCert")
}
