package blockchain

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type FabricClient struct {
	Contract *client.Contract
}

// ChaincodeAPI defines the methods used by the API handlers. FabricClient implements this.
type ChaincodeAPI interface {
	AddInventoryItem(id string, code string, name string, description string, symbol string, conversionFactor int, baseUnit bool, category string, status bool) (string, error)
	ReadItem(id string) (string, error)
	UpdateItem(id string, code string, name string, description string, symbol string, conversionFactor int, baseUnit bool, category string, status bool) (string, error)
	DeleteItem(id string) (string, error)
	// GetHistory returns history JSON string for an id
	GetHistory(id string) (string, error)
}

// ClientConfig contains configuration used to establish a gateway connection.
type ClientConfig struct {
	CertPath      string // path to X.509 cert PEM
	KeyDir        string // directory containing private key file
	TLSCertPath   string // path to peer TLS CA cert
	PeerEndpoint  string // host:port of peer gateway (gRPC)
	MSPID         string // e.g., Org1MSP
	ChannelName   string // e.g., mychannel
	ChaincodeName string // e.g., koperasi
}

// NewFabricClient creates a FabricClient using default test-network paths (backwards compatible).
func NewFabricClient() (*FabricClient, error) {
	return NewFabricClientWithConfig(nil)
}

// NewFabricClientWithConfig creates a FabricClient using the provided configuration.
// If cfg is nil or fields are empty, sensible defaults pointing to the test-network are used.
func NewFabricClientWithConfig(cfg *ClientConfig) (*FabricClient, error) {
	// apply defaults
	if cfg == nil {
		cfg = &ClientConfig{}
	}
	if cfg.CertPath == "" {
		cfg.CertPath = filepath.Join("..", "test-network", "organizations", "peerOrganizations", "org1.example.com", "users", "Admin@org1.example.com", "msp", "signcerts", "cert.pem")
	}
	if cfg.KeyDir == "" {
		cfg.KeyDir = filepath.Join("..", "test-network", "organizations", "peerOrganizations", "org1.example.com", "users", "Admin@org1.example.com", "msp", "keystore")
	}
	if cfg.TLSCertPath == "" {
		cfg.TLSCertPath = filepath.Join("..", "test-network", "organizations", "peerOrganizations", "org1.example.com", "peers", "peer0.org1.example.com", "tls", "ca.crt")
	}
	if cfg.PeerEndpoint == "" {
		cfg.PeerEndpoint = "localhost:7051"
	}
	if cfg.MSPID == "" {
		cfg.MSPID = "Org1MSP"
	}
	if cfg.ChannelName == "" {
		cfg.ChannelName = "mychannel"
	}
	if cfg.ChaincodeName == "" {
		cfg.ChaincodeName = "koperasi"
	}

	// Load identity and signer
	id, sign, err := LoadIdentityFromFiles(cfg.CertPath, cfg.KeyDir, cfg.MSPID)
	if err != nil {
		return nil, err
	}

	// Load TLS Certificate
	caCert, err := ioutil.ReadFile(cfg.TLSCertPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read TLS certificate (%s): %v", cfg.TLSCertPath, err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	creds := credentials.NewClientTLSFromCert(certPool, "")

	conn, err := grpc.Dial(cfg.PeerEndpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to gateway (%s): %v", cfg.PeerEndpoint, err)
	}

	gw, err := client.Connect(id, client.WithSign(sign), client.WithClientConnection(conn))
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to gateway: %v", err)
	}

	network := gw.GetNetwork(cfg.ChannelName)
	contract := network.GetContract(cfg.ChaincodeName)

	return &FabricClient{Contract: contract}, nil
}