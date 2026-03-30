package blockchain

import (
    "fmt"
)

func (fc *FabricClient) AddInventoryItem(id string, code string, name string, description string, symbol string, conversionFactor int, baseUnit bool, category string, status bool) (string, error) {
    result, err := fc.Contract.SubmitTransaction("AddInventoryItem", id, code, name, description, symbol, fmt.Sprintf("%d", conversionFactor), fmt.Sprintf("%t", baseUnit), category, fmt.Sprintf("%t", status))
    
    if err != nil {
        return "", err
    }

    return string(result), nil
}

func (fc *FabricClient) ReadItem(id string) (string, error) {
    result, err := fc.Contract.EvaluateTransaction("ReadItem", id)
    if err != nil {
        return "", err
    }
    return string(result), nil
}

func (fc *FabricClient) UpdateItem(id string, code string, name string, description string, symbol string, conversionFactor int, baseUnit bool, category string, status bool) (string, error) {
    result, err := fc.Contract.SubmitTransaction("UpdateItem", id, code, name, description, symbol, fmt.Sprintf("%d", conversionFactor), fmt.Sprintf("%t", baseUnit), category, fmt.Sprintf("%t", status))
    if err != nil {
        return "", err
    }
    return string(result), nil
}

func (fc *FabricClient) DeleteItem(id string) (string, error) {
    result, err := fc.Contract.SubmitTransaction("DeleteItem", id)
    if err != nil {
        return "", err
    }
    return string(result), nil
}