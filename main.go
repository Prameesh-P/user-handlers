package main

import (
	"database/sql"
	"log"
	"os"
	"user-handler/api"
	dbs "user-handler/db/sqlc"
	"user-handler/initialiazer"

	_ "github.com/lib/pq"
)

// build two api
//  get post
//  psql
//  username and password json
//  get == > send user query params
//  if empty All usersgithub.com/gin-gonic/gin"

func main() {
	initialiazer.LoadEnv()
	dsn := os.Getenv("dsn")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect psql db")
	}
	sqlcDb := dbs.New(db)
	var server api.Server
	server.Db = *sqlcDb
	server.NewServer()
}
