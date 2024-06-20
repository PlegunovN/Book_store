package configs

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

type localConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
	SslMode  string `json:"sslmode"`
}

func LoadConfig(fileName string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath("./configs")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

//1
//db, err := sqlx.Connect("postgres", v.ConfigFileUsed())

//2
//db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
//	v.Get("host"), v.Get("port"), v.Get("user"), v.Get("password"), v.Get("dbname"), v.Get("sslmode")))

//db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=test_books sslmode=disable")
