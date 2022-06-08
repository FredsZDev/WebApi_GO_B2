package main

import (
	"github.com/FredsZDev/WebApi_GO_B2/database"
	"github.com/FredsZDev/WebApi_GO_B2/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()

}
