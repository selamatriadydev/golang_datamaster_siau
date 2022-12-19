package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"backend_api/app/controllers"
	"backend_api/app/database/seed"
	"backend_api/app/database/migration"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	var tb_migrate string
	tb_migrate = os.Getenv("MIGRATE_TABLE")
	if tb_migrate != "" {
		fmt.Println("We are getting the env values", tb_migrate)
		migration.Load(server.DB)
	}
	var tb_seed string
	tb_seed = os.Getenv("SEEDER_TABLE") 
	if tb_seed != ""{
		seed.Load(server.DB)
	}

	server.Run(":8080")

}