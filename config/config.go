package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	Server   Server
	Database Database
}

// Server struct
type Server struct {
	Port string
}

// Database struct
type Database struct {
	URI          string
	DatabaseName string
	Username     string
	Password     string
}

func (c *Config) Read() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	viper.SetConfigType("yml")

	viper.Set("database.username", os.Getenv("MONGO_USER"))
	viper.Set("database.password", os.Getenv("MONGO_PASS"))
	viper.Set("database.databasename", os.Getenv("MONGO_DB"))
	viper.Set("database.uri", os.Getenv("MONGO_URI"))
	viper.Set("server.port", os.Getenv("PORT"))

	if os.Getenv("ENV") == "prod" {
		viper.SetConfigName("config-prod")
	} else {
		viper.SetConfigName("config")
	}
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config, %s", err)
	}
	viper.Unmarshal(&c)
}
