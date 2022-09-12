package domain

// DatabaseConfig contains the configuration for the database connection
type DatabaseConfig struct {
	DBPort     string `mapstructure:"DB_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

// Server port
type ApplicationPort struct {
	Port string `mapstructure:"APPLICATION_PORT"`
}
