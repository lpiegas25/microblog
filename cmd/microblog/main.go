package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/lpiegas25/microblog/internal/data"
	"github.com/lpiegas25/microblog/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	// connection to the database.
	d := data.New()
	errDB := d.DB.Ping()
	if errDB != nil {
		log.Fatal(errDB)
	}

	go serv.Start()

	// Wait for an in interrupt.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	errSrv := serv.Close()
	if errSrv != nil {
		log.Fatal(errSrv)
	}
	errDB = data.Close()
	if errDB != nil {
		log.Fatal(errDB)
	}
}
