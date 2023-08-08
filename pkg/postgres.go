package pkg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Pgdb() (*sqlx.DB, error) {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")

	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, port, dbName)

	return sqlx.Connect("postgres", config)

}
