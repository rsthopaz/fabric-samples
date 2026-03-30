package main

import (
	"fmt"
	"koperasi-app/blockchain"
)

func main() {
	client, err := blockchain.NewFabricClient()
	if err != nil {
		panic(err)
	}

	// ADD ITEM
	res, err := client.AddInventoryItem(
		"100",
		"BOX",
		"Box Karton",
		"Unit packaging",
		"box",
		1,
		false,
		"Quantity",
		true,
	)

	if err != nil {
		panic(err)
	}
	fmt.Println("Add result:", res)

	// READ ITEM
	read, err := client.ReadInventoryItem("100")
	if err != nil {
		panic(err)
	}

	fmt.Println("Read result:", read)
}