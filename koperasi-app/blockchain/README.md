# Blockchain package — documentation

This README explains the `blockchain` package: files, responsibilities, and the flow from the API layer down to the Fabric Gateway.

Location
- `koperasi-app/blockchain`

Purpose
- Encapsulate Fabric Gateway logic and provide typed wrappers for chaincode transactions. Keep Fabric-related code isolated from HTTP handlers to improve testability.

Files
- `client.go`
  - Provides `FabricClient` which wraps a `*client.Contract` from the Fabric Gateway SDK.
  - Exposes `ClientConfig` and `NewFabricClientWithConfig(cfg *ClientConfig)` to configure certificate paths, TLS cert, peer endpoint, MSP ID, channel and chaincode names.
  - Provides a backwards-compatible `NewFabricClient()` that applies sensible test-network defaults.
  - Declares `ChaincodeAPI` interface with methods used by the API handlers (e.g., `AddInventoryItem`, `ReadItem`, `UpdateItem`, `DeleteItem`, `GetHistory`).

- `wallet.go`
  - `LoadIdentityFromFiles(certPath, keyDir, mspID)` reads X.509 certificate PEM and private key PEM and returns an `identity.Identity` and `identity.Sign` for the Fabric Gateway.
  - Centralizes identity loading and error messages.

- `inventory.go`
  - Contains typed wrappers that call `fc.Contract.SubmitTransaction` or `fc.Contract.EvaluateTransaction` for inventory-related functions such as `AddInventoryItem`, `ReadItem`, `UpdateItem`, `DeleteItem`.
  - Converts typed Go arguments to string arguments required by the Gateway.

- `history.go` (if present)
  - Wrapper to call chaincode's `GetHistory` (returns JSON string currently).

Flow (API -> blockchain -> peer)
1. `main.go` constructs a `FabricClient` (via `NewFabricClient()` or `NewFabricClientWithConfig`) and passes it to `api.NewServer`.
2. API handlers call methods on the `ChaincodeAPI` interface (for example `AddInventoryItem`).
3. The `FabricClient` implementation invokes the Fabric Gateway `Contract.SubmitTransaction` for submit operations (write) or `Contract.EvaluateTransaction` for queries (read).
4. The Gateway packages the proposal, sends it to endorsing peers, collects endorsements, and submits the transaction to the orderer.
5. The chaincode executes on peers and returns a payload (if implemented). The Gateway returns bytes to the `FabricClient`.
6. `FabricClient` returns the bytes (as string) to the API handler which forms the HTTP response.

Configuration notes
- Default file paths point to the repository `test-network` layout (relative paths). If your network differs, call `NewFabricClientWithConfig` and set `PeerEndpoint`, `CertPath`, `KeyDir`, and `TLSCertPath` accordingly.

Testing and debugging
- Unit tests: mock the `ChaincodeAPI` interface in API tests. The `blockchain` package itself can be tested against a running test-network in integration tests.
- If you see `connection refused` to `127.0.0.1:7051`, ensure the peer container is running and the endpoint matches `ClientConfig.PeerEndpoint`.


