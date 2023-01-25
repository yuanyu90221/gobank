package main

import (
	"flag"
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

var ApiConfig *Config
var DSN string

func seedAccount(store Storage, firstName, lastName, password string) *Account {
	acc, err := NewAccount(firstName, lastName, password)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	log.Println("new account => ", acc.Number)
	return acc
}
func seedAccounts(s Storage) {
	seedAccount(s, "gson", "liang", "richer123")
}
func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()
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
	// seed stuff
	if *seed {
		log.Println("Seeding the database")
		seedAccounts(store)
	}
	server := NewAPIServer(fmt.Sprintf(":%s", ApiConfig.PORT), store)
	server.Run()
}
