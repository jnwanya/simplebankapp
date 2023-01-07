package main

import (
	"database/sql"
	"jnwanya/simplebank/api"
	db "jnwanya/simplebank/db/sqlc"
	"jnwanya/simplebank/db/util"
	"log"

	_ "github.com/lib/pq"
)

/*
const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:p@ssw0rd@localhost:6543/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)*/

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalln("Cannot load app configs.", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatalln("Cannot connect to the DB.", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatalln("Cannot create server.", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatalln("Cannot start server.", err)
	}
}
