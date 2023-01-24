package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var psqlDSN string

var Client *gorm.DB

func init() {

	host, okHost := os.LookupEnv("NOTES_HOST")
	user, okUser := os.LookupEnv("NOTES_USER")
	pass, okPass := os.LookupEnv("NOTES_PASS")
	dbname, okDbname := os.LookupEnv("NOTES_DBNAME")

	if !okHost || !okUser || !okPass || !okDbname {
		log.Fatal("Ensure correct db config initialization")
		return
	}

	psqlDSN = fmt.Sprintf("host=%v", host) +
		" " + fmt.Sprintf("user=%v", user) +
		" " + fmt.Sprintf("password=%v", pass) +
		" " + fmt.Sprintf("dbname=%v", dbname)
}

func Connect() {
	client, err := gorm.Open(postgres.Open(psqlDSN))

	if err != nil {
		log.Fatal(err)
		return
	}

	Client = client

	log.Println("Database connected")
}
