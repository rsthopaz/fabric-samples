# koperasi-app — REST API for Koperasi Chaincode

This README explains what was implemented in the `koperasi-app` module, why, and how to run and test the HTTP API. It is written for a developer who may be unfamiliar with Hyperledger Fabric and this repo.

## Summary of changes (what I implemented)

- Refactored the Fabric Gateway client to be configurable (see `blockchain/client.go`). You can now create a `FabricClient` using `NewFabricClientWithConfig(cfg)` or the backward-compatible `NewFabricClient()`.
- Added a small `wallet` helper (`blockchain/wallet.go`) to load an X.509 identity and signer from certificate and keystore files.
- Introduced a `ChaincodeAPI` interface in `blockchain/client.go` and implemented it with the existing `FabricClient`. Handlers accept this interface so they can be unit-tested with a mock.
- Scaffolded a Gin-based REST server with inventory handlers:
  - `POST /api/v1/inventory` — add item
  - `GET /api/v1/inventory/:id` — read item
  - `PUT /api/v1/inventory/:id` — update item
  - `DELETE /api/v1/inventory/:id` — delete item
  - `GET /api/v1/inventory/:id/history` — get history
- Added unit tests for the inventory handlers using a mock `ChaincodeAPI` (`api/handlers_inventory_test.go`).

## Important current behavior (short)

- The app uses the Fabric Gateway Go SDK and expects a running Fabric `test-network` with the `koperasi` chaincode deployed.
- When you `POST /api/v1/inventory`, the API currently returns the raw SubmitTransaction result as `tx` — for the current chaincode this is often an empty string because the chaincode does not return a payload.
- `GET /api/v1/inventory/:id` returns the chaincode JSON as a string in the `item` field. This is functional but not prettified; you can request a follow-up patch to return parsed JSON objects.

## Versions / dependencies used here

- Go: tested with Go 1.24.x on Windows (your `go version` output).
- Gin (HTTP): module resolved to `github.com/gin-gonic/gin` (module version used during build: v1.12.0).
- Fabric Gateway SDK: `github.com/hyperledger/fabric-gateway` (version from `go.mod`).

The exact versions are recorded in `koperasi-app/go.mod`.

## Prerequisites (what you need locally)

- Docker and Docker Compose (to run the Fabric test-network scripts).
- A working Go toolchain (Go 1.24+ recommended).
- The repository root (this workspace) containing `test-network` and `koperasi-chaincode` folders.

## Step-by-step: bring up the Fabric test-network and deploy chaincode

1. Open a terminal (use Git Bash or WSL on Windows if recommended by the Fabric scripts).
2. From the repository root (the folder that contains `test-network`), run the network script to start the test network and create the channel:

```bash
cd test-network
./network.sh up createChannel -c mychannel -ca
```

3. Deploy the `koperasi` chaincode (example using the repo scripts):

```bash
# from test-network directory
./network.sh deployCC -ccn koperasi -ccp ../koperasi-chaincode -ccl go
```

4. Verify peer is running and listening on `7051` (host) and that the `Admin` identity and TLS certs are present under `test-network/organizations/...`.

Notes: If your environment or the repo provides different scripts, follow `test-network`'s README.

## Step-by-step: run the API server

1. Build the API server (from `koperasi-app` directory):

```powershell
cd koperasi-app
go build -o koperasi-app.exe main.go
```

2. Run the server (default listen address is `:8080`; in some runs you may see a different port if changed):

```powershell
.\koperasi-app.exe
# OR to run with go run (dev):
go run main.go
```

3. Check the health endpoint:

```http
GET http://localhost:8080/health
```

If you see a 200 JSON `{"status":"ok"}` the HTTP server is up.

## How to test endpoints (example requests)

Use Postman, curl, or any HTTP client.

1. Add an item (POST)

POST http://localhost:8080/api/v1/inventory
Headers: `Content-Type: application/json`
Body (JSON):

```json
{
  "id": "102",
  "code": "BOX",
  "name": "Box",
  "description": "Unit packaging",
  "symbol": "box",
  "conversionFactor": 1,
  "baseUnit": false,
  "category": "Quantity",
  "status": true
}
```

Expected response (current behavior):

```json
{ "tx": "" }
```

Why empty? The `AddInventoryItem` chaincode currently does not return a payload; `SubmitTransaction` returns an empty byte slice. See "Next steps" below to change this behavior.

2. Read an item (GET)

GET http://localhost:8080/api/v1/inventory/102

Current response example:

```json
{
  "item": "{\"ID\":\"102\",\"Code\":\"BOX\",...}"
}
```

Notes: the `item` field contains JSON returned by chaincode as a string. You can parse it client-side, or request a follow-up patch to have the server parse and return proper JSON objects.

3. Update an item (PUT)

PUT http://localhost:8080/api/v1/inventory/102
Body: same shape as POST (id may be ignored since it's in URL)

Response: `{ "tx": "" }` (or tx id if chaincode returns a value or we switch to `SubmitAsync` — see below).

4. Delete an item (DELETE)

DELETE http://localhost:8080/api/v1/inventory/102

Response: `{ "tx": "" }` or an error.

5. History (GET)

GET http://localhost:8080/api/v1/inventory/102/history

Response: `{ "history": "[...]" }` where the `history` is currently the chaincode JSON string.

## Debugging connection errors (e.g., `connection refused`)

- If the API logs show errors like:

```
endorse error: rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing: dial tcp 127.0.0.1:7051: connect: connection refused"
```

then the API cannot reach the Fabric peer at `127.0.0.1:7051`. Steps:

1. Ensure the Fabric test-network is running (see earlier `network.sh up` step).
2. Verify Docker containers: `docker ps` and check `peer0.org1` is running.
3. Confirm TLS certs and Admin identity files are present under `test-network/organizations/...` — the app expects them by default under relative paths.
4. If your peer uses different host/port, update the config when creating the client (use `NewFabricClientWithConfig`) or set up an env-based wrapper (future patch).

## Files you will likely edit next

- `koperasi-app/blockchain/client.go` — client config and connection options.
- `koperasi-app/blockchain/inventory.go` — wrappers to change `SubmitTransaction` to `SubmitAsync` if you want tx IDs.
- `koperasi-app/api/handlers_inventory.go` — change to parse chaincode JSON results into native JSON responses.

## Recommended next improvements (roadmap)

1. Parse chaincode query responses into JSON before returning (improves Postman UX). I can implement this now.
2. Return a transaction ID from submit operations by using `SubmitAsync` and returning the tx ID immediately.
3. Add JWT authentication and role checks to protect modifying endpoints.
4. Add integration test scripts that start `test-network`, deploy chaincode, run the server, and execute example curl/Postman tests.

## Where to look in the code

- Chaincode logic: [koperasi-chaincode/chaincode/smartcontract.go](../koperasi-chaincode/chaincode/smartcontract.go)
- Fabric client: [koperasi-app/blockchain/client.go](blockchain/client.go)
- Wallet loader: [koperasi-app/blockchain/wallet.go](blockchain/wallet.go)
- Inventory wrappers: [koperasi-app/blockchain/inventory.go](blockchain/inventory.go)
- API server: [koperasi-app/api/server.go](api/server.go)
- Inventory handlers: [koperasi-app/api/handlers_inventory.go](api/handlers_inventory.go)
- Unit tests: [koperasi-app/api/handlers_inventory_test.go](api/handlers_inventory_test.go)
