package main

import (
	"biFebriansyah/gogin/internal/routers"
	"biFebriansyah/gogin/pkg"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
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
