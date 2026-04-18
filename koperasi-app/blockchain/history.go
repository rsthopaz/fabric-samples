package blockchain

func (fc *FabricClient) GetHistory(id string) (string, error) {
    result, err := fc.Contract.EvaluateTransaction("GetHistory", id)
    if err != nil {
        return "", err
    }
    return string(result), nil
}