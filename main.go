package main

import (
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

var ApiConfig *Config
var DSN string

func main() {
	// load config
	ApiConfig = LoadConfig()
	DSN = GetDSN(ApiConfig)
	store, err := NewPostgresStore(DSN)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(fmt.Sprintf(":%s", ApiConfig.PORT), store)
	server.Run()
}
