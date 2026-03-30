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
		"101",
		"BOX",
		"Box Karton 101",
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
	read, err := client.ReadItem("100")
	if err != nil {
		panic(err)
	}

	fmt.Println("Read result:", read)

    // UPDATE ITEM
    update, err := client.UpdateItem(
        "101",
        "BOX",
        "Box Karton 101 Updated",
        "Unit packaging updated",
        "box",
        2,
        false,
        "Quantity",
        true,
    )
    if err != nil {
        panic(err)
    }
    fmt.Println("Update result:", update)

    // DELETE ITEM
    delete, err := client.DeleteItem("101")
    if err != nil {
        panic(err)
    }
    fmt.Println("Delete result:", delete)

}