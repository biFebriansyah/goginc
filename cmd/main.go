package main

import (
	"biFebriansyah/gogin/internal/routers"
	"biFebriansyah/gogin/pkg"
	"flag"
	"fmt"
	"log"

	"github.com/asaskevich/govalidator"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	viper.SetConfigName("env.dev")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	serverFlag := flag.Bool("listen", false, "Run function a (server)")
	migrateUpFlag := flag.Bool("migrate-up", false, "Run function to apply migration")
	migrateDownFlag := flag.Bool("migrate-down", false, "Run function to rollback migration")

	flag.Parse()
	migrate := pkg.NewMigrator()

	if *serverFlag {
		listen()
	} else if *migrateUpFlag {
		migrate.Ups()
	} else if *migrateDownFlag {
		migrate.Downs()
	} else {
		fmt.Println("Usage: go run main.go --listen or go run main.go --migrate-up")
	}
}

func listen() {
	database, err := pkg.Pgdb()
	if err != nil {
		log.Fatal(err)
	}
	router := routers.New(database)
	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
