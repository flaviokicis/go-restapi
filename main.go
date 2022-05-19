package main

import (
	"github.com/flaviokicis/go-restapi/database"
	"github.com/flaviokicis/go-restapi/server"
)

// Start database

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
