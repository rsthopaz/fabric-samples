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
		"105",
		"BOX",
		"Box Karton 105",
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
	read, err := client.ReadItem("105")
	if err != nil {
		panic(err)
	}

	fmt.Println("Read result:", read)

    // UPDATE ITEM
    _, err = client.UpdateItem(
        "105",
        "BOX",
        "Box Karton 105 Updated",
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

    updated, err:= client.ReadItem("105")
    if err != nil {
        panic(err)
    }
    fmt.Println("Update result:", updated)

    // DELETE ITEM
    _, err = client.DeleteItem("105")
    if err != nil {
        panic(err)
    }
    
    deleted, err := client.ReadItem("105")
    if err != nil {
        fmt.Println("Item successfully deleted, cannot read:", err)
    } 

    fmt.Println("Delete result:", deleted)

}