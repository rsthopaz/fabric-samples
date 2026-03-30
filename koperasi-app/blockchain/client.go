package blockchain

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/hyperledger/fabric-gateway/pkg/client" 
	"github.com/hyperledger/fabric-gateway/pkg/identity" 
	"google.golang.org/grpc" 
	"google.golang.org/grpc/credentials"
)

type FabricClient struct {
	Contract *client.Contract
}

func NewFabricClient() (*FabricClient, error) {
	// Load the connection profile;
	certPath := filepath.Join("..", "wallet", "admin-msp", "msp", "signcerts", "cert.pem")
	keyDir := filepath.Join("..", "wallet", "admin-msp", "msp", "keystore")
	tlscertPath := filepath.Join("..", "..", "fabric-samples", "test-network", "organizations", "peerOrganizations", "org1.example.com", "peers", "peer0.org1.example.com", "tls", "ca.crt")

	// Load Certificate
	certPem, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read certificate: %v", err)
	}

	cert, err := identity.CertificateFromPEM(certPem)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse certificate: %v", err)
	}

	id, err := identity.NewX509Identity("Org1MSP", cert)
	if err != nil {
		return nil, fmt.Errorf("Failed to create identity: %v", err)
	}

	// Load Private Key
	files, _ := ioutil.ReadDir(keyDir)
	keyPath := filepath.Join(keyDir, files[0].Name())

	keyPem, err := ioutil.ReadFile(keyPath)
	privateKey, err := identity.PrivateKeyFromPEM(keyPem)

	sign, err := identity.NewPrivateKeySign(privateKey)

	// Load TLS Certificate
	caCert, err := ioutil.ReadFile(tlscertPath)
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	creds := credentials.NewClientTLSFromCert(certPool, "")

	conn, err := grpc.Dial("localhost:7051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}

	gw, err := client.Connect(id, client.WithSign(sign), client.WithClientConnection(conn))
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}

	network := gw.GetNetwork("mychannel")
	contract := network.GetContract("koperasi-chaincode")

	return &FabricClient{Contract: contract}, nil
}