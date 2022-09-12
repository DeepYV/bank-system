package config

import (
	domain "github.com/banking/domain"
	"github.com/spf13/viper"
)

// Databaseconfig domain configuration for database connections
var Databaseconfig *domain.DatabaseConfig
var ApplicationPort *domain.ApplicationPort

// GetDatabaseConfig store .json value in struct
func GetDatabaseConfig() error {

	viper.SetConfigFile(`config.yaml`)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Databaseconfig)
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&ApplicationPort)
	if err != nil {
		return err
	}
	return err
}
