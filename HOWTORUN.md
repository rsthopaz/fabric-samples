
# FABRIC

This is only setup for testing network
```
./network.sh down

docker rm -f $(docker ps -aq)
docker volume prune -f
docker network prune -f
docker volume rm $(docker volume ls -q | grep compose_)

./network.sh up createChannel -c mychannel -ca
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go -ccl go


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

peer chaincode invoke \
  -o localhost:7050 \
  --ordererTLSHostnameOverride orderer.example.com \
  --tls \
  --cafile $ORDERER_CA \
  -C mychannel \
  -n basic \
  --waitForEvent \
  --peerAddresses localhost:7051 \
  --tlsRootCertFiles $PEER0_ORG1_CA \
  --peerAddresses localhost:9051 \
  --tlsRootCertFiles $PEER0_ORG2_CA \
  -c '{"Args":["CreateAsset","asset1","blue","5","Tom","300"]}'

peer chaincode query \
  -C mychannel \
  -n basic \
  -c '{"Args":["ReadAsset","asset1"]}'

```

If anything happenes like this one:

```
Error: failed to endorse proposal: rpc error: code = Unknown desc = error validating proposal: access denied: channel [mychannel] creator org unknown, creator is malformed
Chaincode definition approved on peer0.org1 on channel 'mychannel' failed
Deploying chaincode failed
(base) hohohihe@LAPTOP-U5VK2MHE:/mnt/d/Thopaz/Kuliah/Project_BC-Koperasi/fabric-samples/test-network$
```

It means reset all progress or do this instead:

```
./network.sh down

docker rm -f $(docker ps -aq)
docker volume prune -f
docker network prune -f
docker volume rm $(docker volume ls -q | grep compose_)
```

