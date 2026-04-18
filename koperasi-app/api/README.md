# API package — documentation

This document explains the purpose of the `api` package, describes each file, and walks through the runtime flow when handling HTTP requests.

Location
- `koperasi-app/api`

Purpose
- Provide an HTTP JSON REST surface for the `koperasi` chaincode using Gin.
- Translate HTTP requests into calls to the blockchain wrapper (via the `ChaincodeAPI` interface).

Files
- `server.go`
  - Creates a `gin.Engine` and registers routes under `/api/v1`.
  - Exposes `NewServer(fc ChaincodeAPI) *Server` to inject the blockchain implementation (dependency injection for tests).
  - Routes registered: `/api/v1/inventory` (POST), `/api/v1/inventory/:id` (GET, PUT, DELETE) and `/api/v1/inventory/:id/history` (GET), plus `/health`.
  - `Server` holds a `Router *gin.Engine` and `Fc ChaincodeAPI`.

- `handlers_inventory.go`
  - Implements request handlers for inventory endpoints.
  - `InventoryRequest` maps incoming JSON to Go fields.
  - Handlers call `s.Fc.AddInventoryItem`, `s.Fc.ReadItem`, `s.Fc.UpdateItem`, `s.Fc.DeleteItem`, `s.Fc.GetHistory`.
  - Error handling: handlers return JSON `{ "error": "..." }` on failure and appropriate HTTP status codes.

- `handlers_inventory_test.go`
  - Unit tests for the inventory handlers using a `mockChaincode` that implements `ChaincodeAPI`.
  - Tests exercise POST/GET/PUT/DELETE/HISTORY handlers via `httptest` server requests.

Flow (request -> chaincode)
1. Client (Postman/browser) sends HTTP request to API endpoint (e.g., POST `/api/v1/inventory`).
2. Gin router dispatches to the matching handler in `handlers_inventory.go`.
3. Handler validates and binds incoming JSON to `InventoryRequest`.
4. Handler calls the corresponding method on `s.Fc` (type `ChaincodeAPI`): e.g., `AddInventoryItem(...)`.
5. `s.Fc` is implemented by `blockchain.FabricClient` which uses the Fabric Gateway SDK to submit or evaluate transactions.
6. The handler returns the result as JSON to the HTTP client.

Testing
- Unit tests: run `go test ./api` (or `go test ./...` from `koperasi-app`).
- Integration tests: start the Fabric `test-network`, deploy `koperasi` chaincode, then run the API and use curl/Postman to exercise endpoints.

Notes / improvements
- Query responses currently return chaincode JSON as string; consider unmarshalling into native JSON before returning.
- Submit operations currently return raw submit bytes (often empty); consider switching to SubmitAsync or returning chaincode payloads.
