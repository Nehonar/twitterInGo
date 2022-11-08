package main

import (
	"log"

	"github.com/nehonar/twitteringo/db"
	"github.com/nehonar/twitteringo/handlers"
)

func main() {
	if db.CheckConnectionWithPing() == db.StatusConnectionWrong {
		log.Fatal("Whitout connection to MongoDB")
		return
	}
	handlers.HandelersRoute()
}
