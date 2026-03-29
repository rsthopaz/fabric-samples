This is Documentation for Blockchain Koperasi
```
  "data": [
        {
            "id": 1,
            "code": "KG",
            "name": "Kilogram",
            "description": "Unit of mass",
            "symbol": "kg",
            "conversion_factor": 1,
            "base_unit": true,
            "category": "Weight",
            "status": true
        },
```

How to Add using AddInventoryItem
```
peer chaincode invoke -C mychannel -n koperasi \
-c '{"function":"AddInventoryItem","Args":[
  "item1",
  "BRG001",
  "Beras",
  "Beras premium kualitas bagus",
  "kg",
  "1",
  "true",
  "Sembako",
  "true"
]}'
```