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

Before adding or using function from smartcontract.go make sure to follow this step:

```
./network.sh down

docker rm -f $(docker ps -aq)
docker volume prune -f
docker network prune -f
docker volume rm $(docker volume ls -q | grep compose_)

./network.sh up createChannel -c mychannel -ca
./network.sh deployCC -ccn koperasi -ccp ../koperasi-chaincode -ccl go

export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
export PEER0_ORG1_CA=${PWD}/organizations/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem
export PEER0_ORG2_CA=${PWD}/organizations/peerOrganizations/org2.example.com/tlsca/tlsca.org2.example.com-cert.pem
export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem


```

How to Add using AddInventoryItem

```
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile $ORDERER_CA -C mychannel -n koperasi --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA -c '{
  "Args":[
    "AddInventoryItem",
    "10",
    "BOX",
    "Box/Karton",
    "Unit of packaging",
    "box",
    "1",
    "false",
    "Quantity",
    "true"
  ]
}'
```

How to Read using ReadItem
```
peer chaincode query -C mychannel -n koperasi -c '{"Args":["ReadItem","1"]}'
```

How to Update using UpdateItem

How to Delete using DeleteItem
```
peer chaincode query -C mychannel -n koperasi -c '{"Args":["DeleteItem","1"]}'

```