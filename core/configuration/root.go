// Package cmd is part of Gandalf
package configuration

import (
	"github.com/ditrit/gandalf/core/configuration/config"
	"github.com/spf13/viper"
)

var rootCfg = config.NewConfigCmd(
	"gandalf",

	"Gandalf is a tool to easily assemble DevOps software factories",

	`Gandalf stands for Gandalf is A Natural Devops Application Life-cycle Framework.
	Gandalf components and multi language abstract workflow primitives allow you to build or modify in few minutes a DevOps software factory in an efficient, highly secured, enterprise grade way.
	Gandalf philosophy is not to replace or to be a additional layer on existing tools. It only provides a way to easily assemble tools and make them efficiently communicate.`,

	func(cfg *config.ConfigCmd, args []string) {
		mode := viper.GetString("mode")
		cfg.CallSubRun(mode)
	})

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCfg.Execute()
}

func init() {
	// cobra.OnInitialize(initConfig)
	rootCfg.Initialize()

	rootCfg.Key("offset", config.IsInt, "", "Offset used in case of multiple Gandals instances hosted on the same host")

	// flags common to all commands
	rootCfg.Key("lname", config.IsStr, "l", "logical name (non empty value required)")
	rootCfg.SetCheck("lname", CheckNotEmpty)
	rootCfg.SetRequired("lname")
	rootCfg.SetNormalize("lname", TrimToLower)

	rootCfg.Key("config", config.IsStr, "", "config file")

	rootCfg.Key("secret", config.IsStr, "", "Path to the secret (absolute or relative to the configuration directory)")
	rootCfg.SetCheck("secret", CheckNotEmpty)

	rootCfg.Key("max_timeout", config.IsInt, "", "maximum timeout of the connector")
	rootCfg.SetDefault("max_timeout", 100)

	rootCfg.Key("bind", config.IsStr, "", "Address to bind (default is 127.0.0.1)")
	rootCfg.SetDefault("bind", "127.0.0.1")
	rootCfg.SetNormalize("bind", TrimToLower)

	rootCfg.Key("port", config.IsInt, "", "Address to bind (default is *:9099)")
	rootCfg.SetDefault("port", 9099+config.GetOffset())

	rootCfg.Key("cert_dir", config.IsStr, "", "path of the certificates directory (absolute or relative to the configuration directory)")
	rootCfg.SetDefault("cert_dir", "certs")

	rootCfg.Key("cert_pem", config.IsStr, "", "path of the TLS certificate (absolute or relative to the certificates directory)")
	rootCfg.SetDefault("cert_pem", "cert.pem")

	rootCfg.Key("key_pem", config.IsStr, "", "path of the TLS private key (absolute or relative to the certificates directory)")
	rootCfg.SetDefault("key_pem", "key.pem")

	rootCfg.Key("ca_cert_pem", config.IsStr, "", "path of the CA certificate (absolute or relative to the certificates directory)")
	rootCfg.SetDefault("ca_cert_pem", "ca_cert.pem")

	rootCfg.Key("ca_key_pem", config.IsStr, "", "path of the CA key (absolute or relative to the certificates directory)")
	rootCfg.SetDefault("ca_key_pem", "ca_key.pem")

	rootCfg.Key("log_dir", config.IsStr, "", "directory to store gandalf logfile")
	//rootCfg.SetDefault("log_dir", "/var/log/")
	rootCfg.SetDefault("log_dir", "home/romainfairant/gandalf/log/")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	config.InitConfig(rootCfg)
}