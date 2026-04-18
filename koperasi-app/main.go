package main

import (
	"koperasi-app/api"
	"koperasi-app/blockchain"
	"log"
)

func main() {
    client, err := blockchain.NewFabricClient()
    if err != nil {
        log.Fatalf("failed to create fabric client: %v", err)
    }

    srv := api.NewServer(client)
    if err := srv.Run(":50000"); err != nil {
        log.Fatalf("server exited: %v", err)
    }
}