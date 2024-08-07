package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type Postgres struct {
	Host     string `mapstructure:"DB_host"`
	Port     string `mapstructure:"DB_port"`
	User     string `mapstructure:"DB_user"`
	Password string `mapstructure:"DB_password"`
	DbName   string `mapstructure:"DB_name"`
	SslMode  string `mapstructure:"DB_sslmode"`
}

func LoadConfig(path string) (cfg *Postgres, err error) {

	cfg = new(Postgres)

	viper.SetConfigFile(path)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("Read config error %w", err)
		return
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		fmt.Errorf("Unmarshal config error %w", err)
		return
	}

	return cfg, nil
}
