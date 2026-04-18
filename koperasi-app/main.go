package main

import (
    "log"
    "koperasi-app/api"
    "koperasi-app/blockchain"
)

func main() {
    client, err := blockchain.NewFabricClient()
    if err != nil {
        log.Fatalf("failed to create fabric client: %v", err)
    }

    srv := api.NewServer(client)
    if err := srv.Run(":8080"); err != nil {
        log.Fatalf("server exited: %v", err)
    }
}