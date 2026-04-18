package blockchain

import (
    "fmt"
    "io/ioutil"
    "path/filepath"

    "github.com/hyperledger/fabric-gateway/pkg/identity"
)

// LoadIdentityFromFiles reads certificate PEM and private key PEM from the
// provided paths and returns an identity and signer suitable for the Fabric Gateway.
func LoadIdentityFromFiles(certPath string, keyDir string, mspID string) (identity.Identity, identity.Sign, error) {
    certPem, err := ioutil.ReadFile(certPath)
    if err != nil {
        return nil, nil, fmt.Errorf("failed to read certificate (%s): %v", certPath, err)
    }

    cert, err := identity.CertificateFromPEM(certPem)
    if err != nil {
        return nil, nil, fmt.Errorf("failed to parse certificate: %v", err)
    }

    id, err := identity.NewX509Identity(mspID, cert)
    if err != nil {
        return nil, nil, fmt.Errorf("failed to create X509 identity: %v", err)
    }

    files, err := ioutil.ReadDir(keyDir)
    if err != nil || len(files) == 0 {
        return nil, nil, fmt.Errorf("failed to read private key directory (%s): %v", keyDir, err)
    }
    keyPath := filepath.Join(keyDir, files[0].Name())

    keyPem, err := ioutil.ReadFile(keyPath)
    if err != nil {
        return nil, nil, fmt.Errorf("failed to read private key (%s): %v", keyPath, err)
    }

    privateKey, err := identity.PrivateKeyFromPEM(keyPem)
    if err != nil {
        return nil, nil, fmt.Errorf("failed to parse private key: %v", err)
    }

    sign, err := identity.NewPrivateKeySign(privateKey)
    if err != nil {
        return nil, nil, fmt.Errorf("failed to create signer: %v", err)
    }

    return id, sign, nil
}
