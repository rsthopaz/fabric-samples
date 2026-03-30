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
		"104",
		"BOX",
		"Box Karton 104",
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
    _, err = client.UpdateItem(
        "104",
        "BOX",
        "Box Karton 104 Updated",
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

    updated, err:= client.ReadItem("104")
    if err != nil {
        panic(err)
    }
    fmt.Println("Update result:", updated)

    // DELETE ITEM
    _, err = client.DeleteItem("104")
    if err != nil {
        panic(err)
    }
    
    deleted, err := client.ReadItem("104")
    if err != nil {
        fmt.Println("Item successfully deleted, cannot read:", err)
    } 

    fmt.Println("Delete result:", deleted)

}