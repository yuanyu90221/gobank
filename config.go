package main

import (
	"fmt"
	"os"
)

const (
	DB_HOST   = "DB_HOST"
	DB_NAME   = "DB_NAME"
	DB_PORT   = "DB_PORT"
	DB_USER   = "DB_USER"
	DB_PASSWD = "DB_PASSWD"
	PORT      = "PORT"
)

type Config struct {
	DB_HOST   string `json:"db_host"`
	DB_NAME   string `json:"db_name"`
	DB_PORT   string `json:"db_port"`
	DB_USER   string `json:"db_user"`
	DB_PASSWD string `json:"db_passwd"`
	PORT      string `json:"port"`
}

func LoadConfig() *Config {
	config := new(Config)
	config.DB_HOST = os.Getenv(DB_HOST)
	config.DB_NAME = os.Getenv(DB_NAME)
	config.DB_USER = os.Getenv(DB_USER)
	config.DB_PASSWD = os.Getenv(DB_PASSWD)
	config.DB_PORT = os.Getenv(DB_PORT)
	config.PORT = os.Getenv(PORT)
	return config
}

func GetDSN(config *Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
		config.DB_HOST, config.DB_USER, config.DB_PASSWD, config.DB_NAME, config.DB_PORT,
	)
}
