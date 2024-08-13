package main

import (
	"database/sql"
	"log"

	"github.com/karan0033/bank/api"
	db "github.com/karan0033/bank/db/sqlc"
	"github.com/karan0033/bank/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Not able to make a connection")
	}

	store := db.NewStore(conn)
	server, err := api.Newserver(config, store)

	if err != nil {
		log.Fatal("Server is not started:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Server is not started:", err)
	}
}
