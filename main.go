package main

import (
	"log"
	"net/http"
	"time"

	"github.com/MarioGN/desafio-client-server-api/client"
	"github.com/MarioGN/desafio-client-server-api/server"
)

func main() {
	client := client.NewClient()
	srv := server.NewUSDBRLService()

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", srv.USDBRLHandler)

	go func() {
		log.Print("starting server")
		http.ListenAndServe(":8080", mux)
	}()

	for {
		log.Print("starting processing")
		client.Execute()
		log.Print("finishing processing\n")
		time.Sleep(10 * time.Second)
	}
}
