package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

// mockChaincode implements the blockchain.ChaincodeAPI for tests.
type mockChaincode struct{}

func (m *mockChaincode) AddInventoryItem(id, code, name, description, symbol string, conversionFactor int, baseUnit bool, category string, status bool) (string, error) {
    return "tx-add-123", nil
}
func (m *mockChaincode) ReadItem(id string) (string, error) {
    return `{"id":"` + id + `","name":"Test Item"}` , nil
}
func (m *mockChaincode) UpdateItem(id, code, name, description, symbol string, conversionFactor int, baseUnit bool, category string, status bool) (string, error) {
    return "tx-update-123", nil
}
func (m *mockChaincode) DeleteItem(id string) (string, error) {
    return "tx-delete-123", nil
}
func (m *mockChaincode) GetHistory(id string) (string, error) {
    return `[{"tx":"tx1"},{"tx":"tx2"}]`, nil
}

func setupRouter() *Server {
    g := gin.New()
    // use the NewServer to construct routes
    s := NewServer(&mockChaincode{})
    // replace router with gin.New() to avoid default middleware in tests
    s.Router = g
    api := g.Group("/api/v1")
    inv := api.Group("/inventory")
    inv.POST("", s.AddInventory)
    inv.GET(":id", s.GetInventory)
    inv.PUT(":id", s.UpdateInventory)
    inv.DELETE(":id", s.DeleteInventory)
    inv.GET(":id/history", s.GetInventoryHistory)
    return s
}

func TestAddInventory(t *testing.T) {
    s := setupRouter()
    body := `{"id":"100","code":"BOX","name":"Box","conversionFactor":1}`
    req := httptest.NewRequest(http.MethodPost, "/api/v1/inventory", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    s.Router.ServeHTTP(w, req)
    if w.Code != http.StatusCreated {
        t.Fatalf("expected 201, got %d, body: %s", w.Code, w.Body.String())
    }
}

func TestGetInventory(t *testing.T) {
    s := setupRouter()
    req := httptest.NewRequest(http.MethodGet, "/api/v1/inventory/100", nil)
    w := httptest.NewRecorder()
    s.Router.ServeHTTP(w, req)
    if w.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", w.Code)
    }
}

func TestUpdateInventory(t *testing.T) {
    s := setupRouter()
    body := `{"id":"100","code":"BOX","name":"Updated"}`
    req := httptest.NewRequest(http.MethodPut, "/api/v1/inventory/100", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    s.Router.ServeHTTP(w, req)
    if w.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d, body: %s", w.Code, w.Body.String())
    }
}

func TestDeleteInventory(t *testing.T) {
    s := setupRouter()
    req := httptest.NewRequest(http.MethodDelete, "/api/v1/inventory/100", nil)
    w := httptest.NewRecorder()
    s.Router.ServeHTTP(w, req)
    if w.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d, body: %s", w.Code, w.Body.String())
    }
}

func TestGetInventoryHistory(t *testing.T) {
    s := setupRouter()
    req := httptest.NewRequest(http.MethodGet, "/api/v1/inventory/100/history", nil)
    w := httptest.NewRecorder()
    s.Router.ServeHTTP(w, req)
    if w.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d, body: %s", w.Code, w.Body.String())
    }
}
