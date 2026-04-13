# Koperasi API Mapping

Purpose
- Capture the API contract and endpoint mapping to be implemented for the Koperasi application (backend and blockchain interactions).

Scope
- REST/HTTP API surface used by frontends and gateway services.
- Chaincode / blockchain invocation endpoints used by backend services.
- Request/response shapes, status codes, authentication, and example payloads.

Base URL
- Example: `https://api.example.com/v1`

Authentication
- Prefer `Bearer` token (JWT) for HTTP endpoints.
- Mutual TLS or client certs may be required for chaincode/peer administrative calls.

Common response envelope
- Success (200/201):
  {
    "success": true,
    "data": { ... },
    "error": null
  }
- Error (4xx/5xx):
  {
    "success": false,
    "data": null,
    "error": { "code": "INVALID_INPUT", "message": "..." }
  }

Resources (high level)
1. Members
- GET /members
  - Query: `?page=1&limit=50`
  - Response: list of members
- GET /members/{id}
- POST /members
  - Request: { "name","email","idNumber",... }
  - Response: created member

2. Accounts (savings, shares)
- GET /accounts?memberId={id}
- POST /accounts
  - Request: { "memberId","type":"savings" }

3. Loans
- POST /loans/apply
  - Request: { "memberId","amount","termMonths","purpose" }
  - Response: { "loanId","status":"pending" }
- GET /loans/{id}
- POST /loans/{id}/approve
- POST /loans/{id}/repay
  - Request: { "amount","paymentDate" }

4. Transactions
- POST /transactions
  - Request: { "fromAccountId","toAccountId","amount","type" }
- GET /transactions?accountId={id}

5. Assets / Inventory (if applicable)
- CRUD endpoints for cooperatives' physical assets

6. Audit / Ledger Queries (read-only)
- GET /ledger/transactions?startDate=&endDate=&memberId=
- These endpoints may return blockchain-backed proofs or transaction IDs.

Blockchain / Chaincode interaction endpoints
- POST /chaincode/invoke
  - Request: { "chaincode":"koperasi_cc","fcn":"CreateMember","args": [ ... ] }
  - Response: { "txId":"...","status":"submitted" }
- POST /chaincode/query
  - Request: { "chaincode":"koperasi_cc","fcn":"GetMember","args": [ "memberId" ] }
  - Response: query result

API design notes / conventions
- Use HTTP verbs: GET (read), POST (create/commands), PUT/PATCH (update), DELETE (remove)
- Keep command endpoints (state-changing) idempotent where possible and return transaction IDs.
- Use 201 for created resources, 200 for successful reads/commands that return data, 202 for accepted async operations.
- Validation errors: 400 + structured error code.
- Authorization errors: 403. Authentication failures: 401.

Payload schemas (examples)
- Member create request
  {
    "name": "Siti Aminah",
    "email": "siti@example.com",
    "phone": "08123456789",
    "idNumber": "ID123456",
    "address": "..."
  }

- Loan apply request
  {
    "memberId": "member-abc-123",
    "amount": 5000000,
    "termMonths": 12,
    "purpose": "Modal usaha"
  }

Wire / integration points
- Frontend -> API gateway (REST)
- API gateway -> backend services
- Backend -> chaincode via Fabric SDK or a dedicated chaincode gateway endpoint

Open questions / next steps
- Finalize auth model (JWT claims mapping to Fabric identities).
- Define exact schema (JSON Schema / OpenAPI) for each endpoint.
- Decide which operations are synchronous vs async (return `txId` + async status polling).

How to use this document
- Use as the living API mapping before producing OpenAPI specs.
- Update with concrete schemas and examples as implementation decisions are made.

README maintained by: Koperasi dev team
