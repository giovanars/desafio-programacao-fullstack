package config

import "github.com/spf13/viper"

type Config struct {
	AppName       string
	AppEnv        string
	HTTPAddress   string
	MySqlAddress  string
	MySqlDataBase string
}

func NewConfig(opts ...ConfigOption) *Config {
	options := configOptions{}

	for _, o := range opts {
		o.apply(&options)
	}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.AddConfigPath(options.configPath)
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	//Defaults values
	viper.SetDefault("HTTP_ADDRESS", "0.0.0.0:8080")

	return &Config{
		AppName:       viper.GetString("APP_NAME"),
		AppEnv:        viper.GetString("APP_ENV"),
		HTTPAddress:   viper.GetString("HTTP_ADDRESS"),
		MySqlAddress:  viper.GetString("MYSQL_ADDRESS"),
		MySqlDataBase: viper.GetString("MYSQL_DATABASE"),
	}

}

type configOptions struct {
	configPath string
}
type ConfigOption interface {
	apply(*configOptions)
}

type configPathOption string

func (c configPathOption) apply(opts *configOptions) {
	opts.configPath = string(c)
}

func WithConfigPath(configPath string) ConfigOption {
	return configPathOption(configPath)
}
